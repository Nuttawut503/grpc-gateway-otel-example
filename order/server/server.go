package server

import (
	"context"
	"grpc-tracing/order/server/customerpb"
	"grpc-tracing/order/server/orderpb"
	"log"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var tr = otel.Tracer("order-example")

type OrderServer struct {
	CustomerSvcConn *grpc.ClientConn

	orderpb.UnimplementedOrderServer
}

func (s *OrderServer) CreateCustomerOrder(ctx context.Context, req *orderpb.CreateCustomerOrderRequest) (*orderpb.CreateCustomerOrderResponse, error) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.String("extra.customer_id", req.CustomerId), attribute.Float64("extra.price", req.Price))
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	))
	res, err := customerpb.NewCustomerClient(s.CustomerSvcConn).
		GetCustomerCreditLimit(ctx, &customerpb.GetCustomerCreditLimitRequest{
			CustomerId: req.CustomerId,
		})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if res.CustomerCreditLimit < req.Price {
		return nil, status.Error(codes.Canceled, "customer_credit_limit is less than the price")
	}
	return &orderpb.CreateCustomerOrderResponse{
		TransactionId: strconv.Itoa(time.Now().Nanosecond()),
	}, nil
}
