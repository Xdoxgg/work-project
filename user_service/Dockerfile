# User Service Dockerfile
FROM golang:1.20 AS builder
WORKDIR /app
COPY user_service.go .
RUN go build -o user_service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user_service .
CMD ["./user_service"]

