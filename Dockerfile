FROM golang:alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./order-service ./cmd/server/main.go
RUN go build -o ./gateway-service ./cmd/client/main.go

FROM alpine:latest AS orders-service
WORKDIR /app
COPY --from=builder /app/order-service .
COPY --from=builder /app/configs/docker.yaml ./configs/docker.yaml
EXPOSE 50051
ENTRYPOINT ["./order-service"]

FROM alpine:latest AS gateway-service
WORKDIR /app
COPY --from=builder /app/gateway-service .
COPY --from=builder /app/configs/docker.yaml ./configs/docker.yaml
EXPOSE 8080
ENTRYPOINT ["./gateway-service"]