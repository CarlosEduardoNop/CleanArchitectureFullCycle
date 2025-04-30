FROM golang:1.23 AS builder

WORKDIR /app
COPY ../.. /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./app/cmd/server/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /app/api ./

CMD ["./api"]