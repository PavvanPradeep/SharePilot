# Dockerfile for client
FROM docker.io/library/golang:1.20 AS builder

WORKDIR /app

# Copy go.mod file
COPY ../go.mod ./

RUN go mod download

# Copy the rest of the application files
COPY . .

RUN go build -o client .

FROM docker.io/library/alpine:latest

WORKDIR /app

COPY --from=builder /app/client /app/client

CMD ["./client"]