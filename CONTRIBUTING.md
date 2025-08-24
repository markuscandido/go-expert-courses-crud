# Guia de Contribuição 👩‍💻👨‍💻

Obrigado por seu interesse em contribuir com o projeto! Este guia irá te ajudar a começar.

## 🏛️ Entendendo a Arquitetura

Este projeto utiliza **Arquitetura Hexagonal (Portas e Adaptadores)**. Antes de contribuir, é importante entender a organização do código:

- **`internal/core`**: O coração da aplicação. Aqui ficam a lógica de negócio, entidades e interfaces (portas).
  - **`domain`**: Entidades e DTOs puros, sem dependências externas.
  - **`port`**: Interfaces que definem os contratos para os casos de uso e repositórios.
  - **`usecase`**: Implementações da lógica de negócio.
- **`internal/application/adapter`**: A implementação das "portas".
  - **`driving`**: Adaptadores que invocam a aplicação (ex: `graphql`, `grpc`).
  - **`driven`**: Adaptadores que são invocados pela aplicação (ex: `storage/postgres`).

Ao adicionar uma nova funcionalidade, você provavelmente irá:
1.  Criar ou alterar um `usecase`.
2.  Implementar a porta do `driven` adapter (ex: um novo método no repositório).
3.  Expor o `usecase` através de um `driving` adapter (ex: um novo resolver no GraphQL).

## 🚀 Primeiros Passos

1.  **Configuração do Ambiente**
    - Siga as instruções do [README.md](./README.md).
    - Certifique-se de que a aplicação está rodando com Docker Compose:
      ```bash
      docker-compose up --build
      ```

2.  **Encontre uma Tarefa**
    - Verifique as [issues abertas](https://github.com/markuscandido/go-expert-courses-crud/issues).
    - Para novas funcionalidades, abra uma issue para discussão.

## 🔄 Fluxo de Desenvolvimento

```bash
# 1. Faça um fork e clone o repositório

# 2. Crie uma branch a partir da `main`
git checkout -b feature/nome-da-feature

# 3. Desenvolva sua feature
# - Siga as convenções de código abaixo
# - Adicione testes para novas funcionalidades
# - Atualize a documentação, se necessário

# 4. Execute testes e linter
go test -v ./...
golangci-lint run

# 5. Formate seu código
go fmt ./...

# 6. Faça o commit seguindo o padrão Conventional Commits
git add .
git commit -m "feat: sua nova feature"

git push origin feature/nome-da-feature
```

## 📝 Padrões do Projeto

### Código
- Siga o [Effective Go](https://golang.org/doc/effective_go.html).
- Use `goimports` para formatação e organização de imports.
- Documente funções e tipos públicos.

### Testes
- Adicione testes unitários para `usecases` e `adapters`.
- Mantenha uma boa cobertura de testes.

### Migrações de Banco de Dados
- Para qualquer alteração no schema, crie um novo arquivo de migração.
- Use a CLI do `golang-migrate` para criar e gerenciar as migrações.

#### Criando uma nova migração
```bash
migrate create -ext sql -dir sql/migrations -seq nome_da_migracao
```

#### Executando as migrações (via Docker)
O `docker-compose.yml` pode ser adaptado para executar as migrações durante o boot da aplicação.

### Commits
Seguimos o [Conventional Commits](https://www.conventionalcommits.org/):
- `feat:` Nova funcionalidade
- `fix:` Correção de bugs
- `docs:` Documentação
- `style:` Formatação
- `refactor:` Refatoração
- `test:` Testes
- `chore:` Tarefas de manutenção

## 🔄 Processo de Pull Request

1.  Atualize sua branch com a `main`.
2.  Certifique-se que todos os testes e o linter estão passando.
3.  Atualize o `CHANGELOG.md` se sua mudança impacta o usuário.
4.  Envie o Pull Request com uma descrição clara do que foi feito.

## 📄 Licença

Contribuições são licenciadas sob [MIT](LICENSE).