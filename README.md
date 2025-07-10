# Sistema de Estoque em Go

## Ferramentas
Vou rodar uma base de dados postgres e o framework gin do go em um container docker.

## Todo
- Api para adicionar, deletar, modificar e listar produtos no estoque ✅

- Criar um logger ✅

- Tornar o atributo available implicito, isto é, não precisar colocá-lo explicitamente

- criar atributo totalValue que vai pegar o preço e multiplicar pela quantidade

- Se produto tiver sido deletado, retornar que foi deletado para o get por produto

- Criar rota para listar os produtos por tipo

- Criar uma rota somente para produtos deletados
        - listar produtos deletados
        - recuperar produtos deletados

- Criar rota que retorna o preço da filial (soma de todos os atributos totalValue)