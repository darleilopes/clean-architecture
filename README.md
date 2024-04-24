Subir serviços:
```shell
docker-compose up
```

Caso não tenha sido feita a migration:
```shell
make migrate
```
[migrate](https://github.com/golang-migrate/migrate)

Para executar o projeto:
```shell
cd cmd/ordersystem
go run main.go wire_gen.go
```

Notas:
- Para testar os endpoints ver path `/api`
- Para testar o grpc favor utilizar a biblioteca evans

[graphql playground](http://localhost:8080/)
```graphql
query queryOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}

mutation addOrder {
  createOrder(input: {id: "aaaaa", Price: 10, Tax:10})
  {
    id
    Price
    FinalPrice
    Tax
  }
}
```