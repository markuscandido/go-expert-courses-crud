# Go Expert - GraphQL API

API GraphQL desenvolvida em Go para o curso Go Expert.

## 🚀 Começando

### Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (opcional, para ambiente de desenvolvimento)
- Git

### Configuração do Ambiente

1. **Clone o repositório**
   ```bash
   git clone <url-do-repositorio>
   cd graphql
   ```

2. **Configure as variáveis de ambiente**
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```env
   PORT=<port>
   DB_USER=<user>
   DB_PASSWORD=<password>
   DB_HOST=<host>
   DB_PORT=<port>
   DB_NAME=<name>
   ```

3. **Inicie o banco de dados (Docker)**
   ```bash
   docker-compose up -d
   ```

4. **Instale as dependências**
   ```bash
   go mod tidy
   ```

5. **Execute a aplicação**
   ```bash
   go run ./cmd/graphql/main.go
   ```

   A aplicação estará disponível em: http://localhost:<port>

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [GraphQL](https://graphql.org/) - API Query Language
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Docker](https://www.docker.com/) - Conteinerização
- [gqlgen](https://gqlgen.com/) - Gerador de código GraphQL para Go

## 📝 Documentação da API

Acesse o Playground GraphQL em: http://localhost:<port>

### Exemplo de Consulta

```graphql
query {
  categories {
    id
    name
    description
  }
}
```

## 📦 Estrutura do Projeto

```
.
├── cmd/graphql/     # Ponto de entrada da aplicação
├── internal/        # Código interno do projeto
│   ├── config/      # Configurações
│   └── database/    # Camada de banco de dados
├── graph/           # Definições do GraphQL
└── schema.graphqls  # Schema GraphQL
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
