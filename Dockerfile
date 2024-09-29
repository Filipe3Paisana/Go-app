# Etapa 1: Construir a aplicação
FROM golang:1.20-alpine AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app


COPY go.mod ./
COPY HelloGO.go ./
COPY RandomNumbers.go ./


RUN go mod tidy


RUN go build -o exercicio1 HelloGO.go
RUN go build -o exercicio2 RandomNumbers.go

# Etapa 2: Executar a aplicação
FROM alpine:latest


WORKDIR /app


COPY --from=builder /app/exercicio1 .
COPY --from=builder /app/exercicio2 .

# Comando padrão para rodar o binário (pode ser alterado para um exercício específico)
CMD ["./exercicio1"]
