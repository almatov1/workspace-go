FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o message-service ./cmd/message/main.go

FROM alpine:latest
COPY --from=builder /app/message-service /app/message-service
COPY .env /app

WORKDIR /app
CMD ["./message-service"]