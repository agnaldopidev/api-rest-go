# API REST em Go

Este projeto é uma API REST simples desenvolvida em Go, com operações para gerenciar usuários. A API possui operaçẽs de CRUD.

## Tecnologias

- Go
- Biblioteca padrão `net/http`

## Estrutura do Projeto
api-rest-go/ ├── main.go # Código principal do servidor e das rotas ├── go.mod # Gerenciamento de dependências └── README.md # Documentação do projeto

## Pré-requisitos

- [Go](https://golang.org/doc/install) instalado na versão 1.22.2 ou superior.

## Configuração e Execução

1. **Clone o repositório** 

   ```bash
   $ git clone git@github.com:agnaldopidev/api-rest-go.git

2. ***Inicialize o módulo Go***
   cd api-rest-go
   ```bash
   $ go mod init api
   
4. ***Execute o servidor***:
   ```bash
   $ go run main.go

6. ***O servidor agora estará disponível em http://localhost:8080***.

## Exemplo:

1. Lista usuarios
   
   ```bash
   $ curl -X GET http://localhost:8080/users

2. Criar usuarios

  ```bash
  $ curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"nome":"Carlos","email":"carlos@example.com"}'
   
