# Planets-b2w

## Utilização

### Variáveis de ambiente
`MONGO_URL` | URL do banco

`PORT` | porta da API

### Iniciar a API
`make run` ou `go run ./cmd/main.go` na raiz do projeto

### Rodar os testes
`make test` ou `go test ./...`

## Endpoints

### GET

Pesquisar todos:

`/planets`

Pesquisar por id:

`/planets/id/d647faa7-c945-443f-9fc9-a06e7196e8fc`

Pesquisar por nome:

`/planets/name/Tatooine`

### POST
Criar um novo planeta, a quantidade de filmes é adicionada conforme a [api](http://swapi.dev/api/planets/) do Star Wars:

`/planets`

body:
`{
    "name":"Tatooine",
    "climate":"arid",
    "terrain":"desert"
}`

### DELETE
Por id:

`/planets/d647faa7-c945-443f-9fc9-a06e7196e8fc`

