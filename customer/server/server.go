package server

import (
	"context"
	"grpc-tracing/customer/server/customerpb"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var tr = otel.Tracer("customer-example")

type Customer struct {
	ID          string
	Name        string
	CreditLimit float64
}

var _customers = []Customer{{ID: "A310", Name: "Teri", CreditLimit: 40000.00}, {ID: "K423", Name: "Mei", CreditLimit: 500}}

type CustomerServer struct {
	customerpb.UnimplementedCustomerServer
}

func (CustomerServer) GetCustomers(ctx context.Context, req *customerpb.GetCustomersRequest) (*customerpb.GetCustomersResponse, error) {
	res := &customerpb.GetCustomersResponse{}
	for _, customer := range _customers {
		res.Customers = append(res.Customers, &customerpb.GetCustomersResponse_Customer{
			CustomerId:   customer.ID,
			CustomerName: customer.Name,
		})
	}
	return res, nil
}

func (CustomerServer) GetCustomerCreditLimit(ctx context.Context, req *customerpb.GetCustomerCreditLimitRequest) (*customerpb.GetCustomerCreditLimitResponse, error) {
	_, span := tr.Start(ctx, "check-credit-limit",
		trace.WithAttributes(attribute.String("extra.customer_id", req.CustomerId)))
	defer span.End()
	for _, customer := range _customers {
		if req.CustomerId == customer.ID {
			return &customerpb.GetCustomerCreditLimitResponse{
				CustomerId:          customer.ID,
				CustomerName:        customer.Name,
				CustomerCreditLimit: customer.CreditLimit,
			}, nil
		}
	}
	return nil, status.Error(codes.NotFound, "customer_id is invalid")
}
