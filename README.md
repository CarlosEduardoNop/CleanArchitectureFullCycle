# Desafio - Listagem de Orders

## Como executar

1. Rode o projeto com:
2. 
```bash
docker compose build
```

```bash
docker compose up -d
```

2. Acesse os serviços:

- REST: [http://localhost:8080/order](http://localhost:8080/order)
- GraphQL: [http://localhost:8081/graphql](http://localhost:8081/graphql)
- gRPC: Porta 50051

## Serviços

| Tipo     | Porta  | Endpoint           |
|----------|--------|--------------------|
| REST     | 8080   | GET or POST /order |
| GraphQL  | 8081   | /graphql           |
| gRPC     | 50051  | ListOrders         |