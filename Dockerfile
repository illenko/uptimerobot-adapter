FROM golang:1.23.0-alpine3.20 AS builder
WORKDIR /build
COPY . .
RUN go build  -o /build/main

FROM ubuntu:24.04
WORKDIR /app
COPY --from=builder /build/main .
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates

EXPOSE 8080
ENTRYPOINT ["/app/main"]