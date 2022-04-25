# gRPC Gateway with Otelgrpc

Modules
- [gRPC | Go](https://github.com/grpc/grpc-go)
- [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [OpenTelemetry-go](https://github.com/open-telemetry/opentelemetry-go)
- [otelgrpc from OpenTelemetry-contrib](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/google.golang.org/grpc/otelgrpc)


This application has two services - Customer and Order
- Customer, I added two default customers and their IDs are "A310" and "K423"
  each customer has credit_limit (will explain about why it exists)
- Order, when making an order, we need a customer_id and the order_price
  if the order_price is more than the customer's credit_limit then the order
  service won't allow the order to be created. But if the order_price is less
  than or equal to the customer's credit_limit then the order service will create
  the order and return transaction_id back to the customer

For my gRPC, it has three functions
- CustomerService; GetCustomers()
  get all customers' info
- CustomerService; GetCustomerCreditLimit(),
  get credit_limit of a customer (customer_id is required)
- OrderService; CreateCustomerOrder()
  to make an order (customer_id and the price is required)
  when a customer request for an order, the service will talk
  to the customer service to get the customer's credit_limit
  then check if it is ok to continue the order. If the order
  service allows, the customer will get a transaction_id in
  response else the customer will get an error and can't make
  the order

How to run
- Docker or Jaeger required
```sh
# if docker
chmod +x jaeger-docker.sh
./jaeger-docker.sh
# open two terminals and run these two services individually
go run grpc-tracing/customer
go run grpc-tracing/order
```

Available URLs to test
- GET localhost:8081/customers
- POST localhost:8082/order

Example
```sh
curl localhost:8081/customers
curl localhost:8082/order \
-H 'Content-Type: application/json' \
-d '{"customerId": "A310", "price": 22}'
```

Visit the Jaeger UI [http://localhost:16686/](http://localhost:16686/)
