FROM docker.io/library/golang:1.20 AS builder

WORKDIR /app


COPY ../go.mod ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM docker.io/library/alpine:latest

WORKDIR /app

COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server"]