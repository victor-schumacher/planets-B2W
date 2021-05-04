# Planets-b2w

## Utilização

### Variáveis de ambiente
`MONGO_URL` | URL do banco

`PORT` | porta da API

### Iniciar a API
`make run` ou `go run ./cmd/main.go` na raiz do projeto

### Rodar os testes
`make test` ou `go test ./...

## Endpoints

### GET

pesquisar todos
/planets

pesquisar por id
/planets/id/d647faa7-c945-443f-9fc9-a06e7196e8fc

pesquisar por nome
/planets/name/Tatooine

### POST
criar um novo planeta, a quantidade de filmes é adicionada conforme a [api](http://swapi.dev/api/planets/) do Star Wars

/planets

body:
`{
    "name":"Tatooine",
    "climate":"arid",
    "terrain":"desert"
}`

### Delete
deletar por id
/planets/d647faa7-c945-443f-9fc9-a06e7196e8fc

