FROM golang:1.23-alpine AS builder

RUN apk --no-cache add

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/app ./main.go

FROM alpine AS runner

COPY --from=builder /app/bin/app /
EXPOSE 8080
CMD ["./app"]
