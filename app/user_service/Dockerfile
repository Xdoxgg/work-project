FROM golang:1.21-alpine AS builder

RUN apk --no-cache add

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/app ./user_service.go


FROM alpine AS runner

COPY --from=builder /app/bin/app /
EXPOSE 8082
CMD ["/app"]