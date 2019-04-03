# Travel Scanner

Travel Scanner é um aplicação para achar o vôo mais barato independente do número de conexões.


## Como Executar

Para que o arquivo seja encontrado dentro do container você deve incluir o mesmo na pasta resources e ao executar o projeto passar o caminho todo.

Caso o arquivo não exista iremos criar um novo que estará disponível nas próximas execuções.

1. Construir imagem docker e executar os testes unitários
```
docker-compose build
```

2. Executar o projeto
**Obs: Para que o arquivo csv seja encontrado na hora de executar o projeto ele deve estar dentro da pasta resources e deve-se usar o caminho /resources/<filename>

Para rest api:
```
docker-compose run -p 8080:8080 api ./api /resources/file.csv
```

Para shell
```
docker-compose run -p 8080:8080 api ./shell /resources/file.csv
```

## Shell

### Consultar um vôo
Para consultar um vôo digitar os códigos dos aeroportos separados por espaço.
```
$ Please enter route: ORI DES
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

## Estrutura

* application: contém as classes de serviço que contém a lógica de negócio
* cmd: contém os arquivos que irão definir como a aplicação irá ser executada
* domain: modelos e interface de repositórios
* errors: mapeamento dos errors da aplição
* infrastructure: recursos relacionados a infraestrutura, como por exemplo, banco de dados
* testing: pacote com todos os teste unitários
* resources: volume a ser mapeado no container docker