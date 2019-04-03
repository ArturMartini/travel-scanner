# Travel Scanner

Travel Scanner é um aplicação para achar o vôo mais barato independente do número de conexões


## Como Executar

Para que o arquivo seja encontrado dentro do container você deve incluir o mesmo na pasta resources e ao executar o projeto passar o caminho todo.

Caso o arquivo não exista iremos criar um novo que estará disponível nas próximas execuções.

1. Construir imagem docker e executar os testes unitários
```
docker-compose build
```

2. Executar o projeto
```
docker-compose run -p 8080:8080 api ./api /resources/file.csv
```


## Rest Api

### Adicionar um novo vôo
**POST http://localhost:8080/flights**
```
{
    "from" : "GRU",
    "to"   : "MIA",
    "cost" : 20
}
```


Exemplo:
```
curl -d '{"from" : "GRU", "to" : "MIA", "cost" : 200}' -H 'Content-Type: application/json' -X POST http://localhost:8080/flights
```

### Procurar pela melhor combinação de vôos
**POST http://localhost:8080/flights/search**
```
{
    "from" : "GRU",
    "to"   : "MIA"
}
```


Exemplo:
```
curl -d '{"from" : "GRU", "to" : "MIA"}' -H 'Content-Type: application/json' -X POST http://localhost:8080/flights/search
```