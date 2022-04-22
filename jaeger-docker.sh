#!/bin/sh

docker run -d \
  -p 6831:6831 \
  -p 14268:14268 \
  -p 16686:16686 \
  jaegertracing/all-in-one
