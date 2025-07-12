# Sistema de Estoque em Go

## Como rodar sistema
- Primeiro você vai criar um arquivo `.env`:
```
PASSWORD=1234
HOST=localhost
PORT=5432
DBNAME=postgres
```
- Depois é só rodar o comando `go run main.go` no terminal

## Requisições
- Endpoint: `/api`
    - GET `/api/products`
    - GET `/api/product/:id`
    - POST `/api/product`
    - PATCH `/api/product/:id`
    - DELETE `/api/product/:id`

### Cadastrar um produto
- É necessário solicitar o nome, tipo, preço e quantidade (opcional)      
