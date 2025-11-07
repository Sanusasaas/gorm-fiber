FROM golang:1.24.8 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN ls -la /app/.env
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/myapp .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .
EXPOSE 5432
CMD ["./myapp"]