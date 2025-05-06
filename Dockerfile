# Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o /motivora-api main.go

# Финальный образ
FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /motivora-api .

EXPOSE 8080

CMD ["./motivora-api"]