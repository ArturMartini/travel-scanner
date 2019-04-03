# Travel Scanner

Travel Scanner é um aplicação para achar o vôo mais barato independente do número de conexões


## Como Executar

Para que o arquivo seja encontrado dentro do container você deve incluir o mesmo na pasta resources e ao executar o projeto passar o caminho todo.

Caso o arquivo não exista iremos criar um novo que estará disponível nas próximas execuções.

1. Construir imagem docker
`
docker-compose build
`

2. Executar o projeto
`
docker-compose run -p 8080:8080 api ./api /resources/file.csv
`
