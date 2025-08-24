# Go Expert - Clean Architecture com GraphQL e gRPC

Este projeto demonstra a implementação de uma aplicação em Go utilizando **Arquitetura Limpa (Clean Architecture)**. A lógica de negócio principal é isolada das camadas de entrega e de persistência, permitindo que o sistema exponha simultaneamente uma API **GraphQL** e um servidor **gRPC**.

A aplicação possui pontos de entrada separados para cada tecnologia, permitindo que os servidores sejam executados de forma independente.

## 🏛️ Arquitetura

O projeto segue os princípios da Arquitetura Limpa, com as seguintes camadas:

- **`internal/entity`**: Contém as entidades de domínio puras, que representam o núcleo do negócio.
- **`internal/usecase`**: Orquestra a lógica de negócio (casos de uso) e define as interfaces dos repositórios.
- **`internal/database`**: Implementa as interfaces de repositório para interagir com o banco de dados PostgreSQL.
- **`graph/`**: Adaptador de entrada que implementa a API GraphQL.
- **`internal/grpc`**: Adaptador de entrada que implementa os serviços gRPC.
- **`internal/web`**: Módulo de inicialização que centraliza a configuração, injeção de dependências e o startup dos servidores.
- **`cmd/`**: Contém os pontos de entrada (`main.go`) para cada servidor.

## 🚀 Começando

### 🛠️ Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Git
- `protoc` (Compilador de Protocol Buffers)

### Configuração do Ambiente

1. **Clone o repositório**
   ```bash
   git clone <url-do-repositorio>
   cd graphql
   ```

2. **Configure as variáveis de ambiente**
   Copie `.env.example` para `.env` e ajuste as configurações do banco de dados, se necessário.
   ```bash
   cp .env.example .env
   ```

3. **Inicie o banco de dados**
   Use o Docker Compose para subir o PostgreSQL.
   ```bash
   docker-compose up -d db
   ```

4. **Instale as dependências**
   ```bash
   go mod tidy
   ```

## 🚀 Executando a Aplicação

Você pode iniciar cada servidor de forma independente.

### Servidor GraphQL

Para iniciar o servidor GraphQL:
```bash
go run cmd/graphql_server/main.go
```
A aplicação estará disponível em `http://localhost:8080`.

### Servidor gRPC

Para iniciar o servidor gRPC:
```bash
go run cmd/grpc_server/main.go
```
O servidor gRPC estará escutando na porta `50051`.

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [GraphQL](https://graphql.org/) - API Query Language
- [gRPC](https://grpc.io/) - Framework de RPC
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Docker](https://www.docker.com/) - Conteinerização
- [gqlgen](https://gqlgen.com/) - Gerador de código GraphQL para Go
- [golang-migrate](https://github.com/golang-migrate/migrate) - Ferramenta de migrações de banco de dados

## 🌐 API REST

A aplicação também expõe uma API REST para gerenciar categorias e cursos.

### Endpoints

#### Categorias
- `POST /categories` - Cria uma nova categoria
- `GET /categories` - Lista todas as categorias
- `GET /categories/{id}` - Obtém uma categoria pelo ID

#### Cursos
- `POST /courses` - Cria um novo curso
- `GET /courses` - Lista todos os cursos
- `GET /courses/{id}` - Obtém um curso pelo ID
- `GET /courses/{id}/category` - Obtém a categoria associada a um curso específico

### Exemplos de Uso

#### Criar uma categoria
```http
POST /categories
Content-Type: application/json

{
  "name": "Programação",
  "description": "Cursos de programação"
}
```

#### Listar todos os cursos
```http
GET /courses
```

#### Obter a categoria de um curso específico
```http
GET /courses/123/category
```

## 📝 Exemplos de Uso da API GraphQL

Acesse o Playground GraphQL em: `http://localhost:8080`

*O conteúdo das queries e mutations permanece o mesmo...*

## 📦 Estrutura do Projeto

```
.
├── cmd/
│   ├── graphql_server/ # Ponto de entrada da aplicação GraphQL
│   └── grpc_server/    # Ponto de entrada da aplicação gRPC
├── internal/
│   ├── config/         # Configurações
│   ├── database/       # Adaptador de banco de dados
│   ├── entity/         # Entidades de domínio
│   ├── grpc/           # Adaptador de serviços gRPC
│   ├── pb/             # Código gerado pelo Protobuf
│   ├── usecase/        # Casos de uso e interfaces
│   └── web/            # Lógica de inicialização dos servidores
├── graph/              # Definições e resolvers do GraphQL
├── proto/              # Arquivos de definição do Protobuf (*.proto)
├── sql/                # Migrações de banco de dados
└── ...
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
