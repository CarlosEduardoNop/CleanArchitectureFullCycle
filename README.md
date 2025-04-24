# Desafio - Listagem de Orders

## Como executar

1. Rode o projeto com:

```bash
docker compose up -d
```

```bash
go run ./cmd/server/main.go
```

2. Acesse os serviços:

- REST: [http://localhost:8080/order](http://localhost:8080/order)
- GraphQL: [http://localhost:8081/graphql](http://localhost:8081/graphql)
- gRPC: Porta 50051

## Serviços

| Tipo     | Porta  | Endpoint                           |
|----------|--------|------------------------------------|
| REST     | 8080   | GET /order                         |
| GraphQL  | 8081   | POST /graphql                      |
| gRPC     | 50051  | ListOrders                         |