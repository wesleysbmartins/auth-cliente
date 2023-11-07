FROM golang:1.20.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 8080

CMD ["/app/main"]
