package main

import (
	"context"
	"grpc-tracing/customer/server/customerpb"
	"grpc-tracing/tracer"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"grpc-tracing/customer/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const (
	_restCustomerAddr = ":8081"
	_grpcCustomerAddr = ":9091"
)

func main() {
	provider, err := tracer.NewTraceProvider("customer-service")
	if err != nil {
		log.Fatal(err)
	}
	defer provider.Shutdown(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := runGRPC(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if customerpb.RegisterCustomerHandlerFromEndpoint(context.Background(), mux, _grpcCustomerAddr, opts) != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:    _restCustomerAddr,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		stop()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Gateway server is running...")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Gateway server is closed...")
}

func runGRPC(ctx context.Context) error {
	lis, err := net.Listen("tcp", _grpcCustomerAddr)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	customerpb.RegisterCustomerServer(srv, &server.CustomerServer{})

	go func() {
		<-ctx.Done()
		srv.Stop()
	}()

	log.Println("gRPC server is running...")
	if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		return err
	}
	log.Println("gRPC server is closed...")
	return nil
}
