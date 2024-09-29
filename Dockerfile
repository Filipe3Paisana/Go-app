# Etapa 1: Construir a aplicação
FROM golang:1.20-alpine AS builder

WORKDIR /app


COPY go.mod ./
COPY main.go ./


RUN go mod tidy

# Compilar o aplicativo Go
RUN go build -o main .

# Etapa 2: Executar a aplicação
FROM alpine:latest


WORKDIR /app


COPY --from=builder /app/main .


CMD ["./main"]
