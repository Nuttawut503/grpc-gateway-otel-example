package main

import (
	"context"
	"grpc-tracing/order/server"
	"grpc-tracing/order/server/orderpb"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"grpc-tracing/tracer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

const (
	_grpcCustomerAddr = ":9091"

	_restOrderAddr = ":8082"
	_grpcOrderAddr = ":9092"
)

func main() {
	provider, err := tracer.NewTraceProvider("order-service")
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
	if orderpb.RegisterOrderHandlerFromEndpoint(context.Background(), mux, _grpcOrderAddr, opts) != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:    _restOrderAddr,
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
	lis, err := net.Listen("tcp", _grpcOrderAddr)
	if err != nil {
		log.Fatalln(err)
	}

	customerConn := mustConnect(_grpcCustomerAddr)
	defer customerConn.Close()

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	orderpb.RegisterOrderServer(srv, &server.OrderServer{
		CustomerSvcConn: customerConn,
	})

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

func mustConnect(port string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		port,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
