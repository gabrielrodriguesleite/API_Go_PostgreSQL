# API_Go_PostgreSQL

## Docker bd PostgreSQL

Subir o container api-todo rodando PostgreSQL.

`docker run -d --name api-todo -p 5432:5423 -e POSTGRES_PASSWORD=1234 postgres:13.5`

**postgres:14.2-alpine utilizada com sucesso(é uma imagem reduzida)**

Executar o cli do PostgreSQL entrando com o usuário padrão.

`docker exec -it api-todo psql -U postgres`

Cria novo usuário.

`create user user_todo;`

Modifica o acesso do usuário.

`alter user user_todo with encrypted password '1122';`

Configura privilégios do usuário para a tabela em questão.

`grant all privileges on database api_todo to user_todo;`

Conectar usando o novo usuário.

Criar o banco de dados.

`create database api_todo;`

Criar a tabela.

`create table todos (id serial primary key, title varchar, description text, done bool default FALSE);`

Permitir o acesso a tabela ao público:

`grant all privileges on all sequences in schema public to user_todo;`

Outros comando úteis:

`\l # para listar as tabelas`

`\c api_todo; # para conectar ao banco em questão`

`\dt # para conferir os dados que por enquanto não existem`


```js
/*
Os passos são os seguintes:
Conectar com o usuário padrão
Criar o usuário user_todo
Configurar a senha e as permissões

Conectar como usuário user_todo
Criar o banco de dados
Criar a tabela
Configurar a tabela como pública

----

Como a senha é definida durante a criação do container é possível 
 configurar uma senha e criar o banco de dados de forma dinâmica com o usuário padrão.

*/
```

Configurar `config.toml` com os mesmos dados de usuário, senha e nome do banco.

Rodar localmente a aplicação com o comando:

`go run main.go`

Testar com postman com um get na porta configurada (9000)

## Docker Compose

Subir o banco com o docker compose e executar a configuração como no exemplo do docker.

`docker-compose up`

Derrubar o container e limpar containers orfãos:

`docker-compose down --remove-orphans`

###### Referências

https://www.youtube.com/watch?v=MD7b-iQMC24

https://www.youtube.com/watch?v=QP0slFxVQKU