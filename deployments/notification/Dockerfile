FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o notification-service ./cmd/notification/main.go

FROM alpine:latest
COPY --from=builder /app/notification-service /app/notification-service
COPY .env /app

WORKDIR /app
CMD ["./notification-service"]