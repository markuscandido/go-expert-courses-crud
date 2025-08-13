# Go Expert - GraphQL API

API GraphQL desenvolvida em Go para o curso Go Expert. Este projeto implementa um sistema de gerenciamento de cursos e categorias, com suporte a operações CRUD e relacionamentos entre entidades.

## 🚀 Começando

### 🛠️ Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose (recomendado para ambiente de desenvolvimento)
- Git
- PostgreSQL 15+ (ou Docker para executar em container)

### Configuração do Ambiente

1. **Clone o repositório**
   ```bash
   git clone <url-do-repositorio>
   cd graphql
   ```

2. **Configure as variáveis de ambiente**
   - Faça uma cópia do arquivo `.env.example` para `.env`:
     ```bash
     cp .env.example .env
     ```
   - Edite o arquivo `.env` com as configurações do seu ambiente

3. **Inicie o banco de dados (Docker)**
   ```bash
   docker-compose up -d
   ```

4. **Instale as dependências**
   ```bash
   go mod tidy
   ```

## 🚀 Executando a Aplicação

### Usando Docker (Recomendado)
```bash
docker-compose up -d
```

### Localmente
1. **Inicie o banco de dados** (caso não esteja usando Docker):
   ```bash
   docker-compose up -d postgres
   ```

2. **Execute as migrações** (já são executadas automaticamente na inicialização):
   ```bash
   go run cmd/graphql/main.go
   ```

3. **Acesse o Playground GraphQL**
   Abra o navegador em: `http://localhost:8080`

## 🛠️ Tecnologias

- [Go](https://golang.org/) - Linguagem de programação
- [GraphQL](https://graphql.org/) - API Query Language
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados
- [Docker](https://www.docker.com/) - Conteinerização
- [gqlgen](https://gqlgen.com/) - Gerador de código GraphQL para Go

## 📝 Documentação da API

Acesse o Playground GraphQL em: `http://localhost:8080`

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
