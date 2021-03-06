# Step Peratama
FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download -x

COPY . .

RUN go build -o main

# Step Kedua
FROM alpine:3.14

WORKDIR /app/webservice

COPY --from=builder /app/.env ./.env
COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"] 
