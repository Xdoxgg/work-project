# Product Service Dockerfile
FROM golang:1.20 AS builder
WORKDIR /app

# Копируем go.mod и go.sum (если есть)
COPY go.mod .
COPY go.sum .

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY product_service.go .

# Собираем исполняемый файл
RUN go build -o product_service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/product_service .
CMD ["./product_service"]