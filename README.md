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
- É necessário solicitar o nome, tipo, quantidade e preço      

## Todo
- Api para adicionar, deletar, modificar e listar produtos no estoque ✅

- Criar um logger ✅

- Tornar o atributo available implicito, isto é, não precisar colocá-lo explicitamente ✅

- criar atributo totalValue que vai pegar o preço e multiplicar pela quantidade ✅

- Se produto tiver sido deletado, retornar que foi deletado para o get por produto

- Criar rota para listar os produtos por tipo

- Criar uma rota somente para produtos deletados
        - listar produtos deletados
        - recuperar produtos deletados

- Criar rota que retorna o preço da filial (soma de todos os atributos totalValue)