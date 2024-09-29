FROM golang:1.20-alpine


WORKDIR /GoApp

COPY go.mod ./
COPY go.sum ./


RUN go mod download


COPY . .


RUN go build -o main .


CMD ["./main"]
