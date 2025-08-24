# Go Expert - Clean Architecture com GraphQL e gRPC

Este projeto demonstra a implementação de uma aplicação em Go utilizando **Arquitetura Hexagonal (Portas e Adaptadores)**. A lógica de negócio principal é isolada das camadas de entrega e de persistência, permitindo que o sistema exponha simultaneamente uma API **GraphQL** e um servidor **gRPC**.

A aplicação possui pontos de entrada separados para cada tecnologia, permitindo que os servidores sejam executados de forma independente ou em conjunto.

## 🏛️ Arquitetura

O projeto segue os princípios da **Arquitetura Hexagonal**, que promove um baixo acoplamento e alta coesão ao isolar o núcleo da aplicação (domínio) de detalhes de implementação externos (adaptadores).

- **Core (Domínio)**: O centro da arquitetura, contendo a lógica de negócio pura.
  - `internal/core/domain`: Entidades, DTOs e erros de domínio.
  - `internal/core/port`: As "Portas", que são interfaces que definem como o mundo exterior pode interagir com o domínio (casos de uso) e como o domínio interage com sistemas externos (repositórios).
  - `internal/core/usecase`: Implementações das portas de entrada (casos de uso), orquestrando a lógica de negócio.

- **Adaptadores**: Componentes que "adaptam" tecnologias externas para as portas da aplicação.
  - **Driving Adapters** (`internal/application/adapter/driving`): Adaptadores que "dirigem" a aplicação, como as APIs.
    - `graphql`: Implementação da API GraphQL.
    - `grpc`: Implementação dos serviços gRPC.
  - **Driven Adapters** (`internal/application/adapter/driven`): Adaptadores que são "dirigidos" pela aplicação, como os clientes de banco de dados.
    - `storage/postgres`: Implementação do repositório para PostgreSQL.
    - `storage/memory`: Implementação do repositório em memória (para testes ou ambientes de desenvolvimento).

- **Aplicação**:
  - `cmd`: Pontos de entrada (`main.go`) para cada servidor (GraphQL e gRPC).
  - `internal/application/server.go`: Configuração e inicialização dos servidores.

## 🚀 Começando

### 🛠️ Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Git
- `protoc` (Compilador de Protocol Buffers)

### Configuração do Ambiente

1. **Clone o repositório**
   ```bash
   git clone https://github.com/markuscandido/go-expert-courses-crud.git
   cd go-expert-courses-crud
   ```

2. **Configure as variáveis de ambiente**
   Copie `.env.example` para `.env`. As configurações padrão já estão prontas para o ambiente Docker.
   ```bash
   cp .env.example .env
   ```

3. **Instale as dependências e gere o código**
   ```bash
   go mod tidy
   go generate ./...
   ```

## 🚀 Executando a Aplicação

### Com Docker Compose (Recomendado)

A forma mais simples de executar a aplicação é com o Docker Compose, que irá iniciar o banco de dados e os servidores GraphQL e gRPC.

```bash
docker-compose up --build
```

- **GraphQL Server**: `http://localhost:8080`
- **gRPC Server**: `localhost:50051`

### Manualmente

Você também pode iniciar cada servidor de forma independente.

1. **Inicie o banco de dados**
   ```bash
   docker-compose up -d postgres
   ```

2. **Execute as migrações do banco de dados**
   Atenção: certifique-se de ter o `golang-migrate` instalado. Veja as instruções no `CONTRIBUTING.md`.
   ```bash
   migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path sql/migrations up
   ```

3. **Inicie os servidores**
   - **Servidor GraphQL**:
     ```bash
     go run cmd/graphql_server/main.go
     ```
   - **Servidor gRPC**:
     ```bash
     go run cmd/grpc_server/main.go
     ```

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [GraphQL](https://graphql.org/) - API Query Language
- [gRPC](https://grpc.io/) - Framework de RPC
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Docker](https://www.docker.com/) - Conteinerização
- [gqlgen](https://gqlgen.com/) - Gerador de código GraphQL para Go
- [golang-migrate](https://github.com/golang-migrate/migrate) - Ferramenta de migrações de banco de dados

## 📝 Exemplos de Uso da API GraphQL

Acesse o Playground GraphQL em: `http://localhost:8080`

*O conteúdo das queries e mutations permanece o mesmo...*

## 📦 Estrutura do Projeto

```
.
├── api/                  # Definições gRPC e Protobuf
├── cmd/                  # Pontos de entrada da aplicação (main.go)
│   ├── graphql_server/
│   └── grpc_server/
├── internal/
│   ├── application/
│   │   ├── adapter/
│   │   │   ├── driven/   # Adaptadores "dirigidos" (banco de dados, etc.)
│   │   │   └── driving/  # Adaptadores "diretores" (GraphQL, gRPC)
│   │   └── server.go     # Configuração e inicialização dos servidores
│   └── core/
│       ├── domain/       # Entidades, DTOs e erros de domínio
│       ├── port/         # Interfaces (Portas)
│       └── usecase/      # Lógica de negócio (Casos de Uso)
├── sql/                  # Migrações de banco de dados
└── ...
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.