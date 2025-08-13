# Guia de Contribuição 👩‍💻👨‍💻

Obrigado por contribuir para o projeto Go Expert GraphQL! Este guia irá te ajudar a começar.

## 🚀 Primeiros Passos

1. **Configuração do Ambiente**
   - Siga as instruções do [README.md](./README.md)
   - Certifique-se de que todos os testes estão passando:
     ```bash
     go test ./...
     ```

2. **Encontre uma Tarefa**
   - Verifique as [issues abertas](https://github.com/markuscandido/go-expert-graphql/issues)
   - Para novas funcionalidades, abra uma issue para discussão

## 🔄 Fluxo de Desenvolvimento

```bash
# 1. Faça um fork e clone o repositório
git clone https://github.com/seu-usuario/go-expert-graphql.git
cd go-expert-graphql

# 2. Crie uma branch
git checkout -b feature/nome-da-feature

# 3. Desenvolva sua feature
# - Siga as convenções de código abaixo
# - Adicione testes para novas funcionalidades
# - Atualize a documentação

# 4. Execute testes e formate o código
golangci-lint run
go test -v ./...

go fmt ./...

# 5. Faça o commit seguindo o padrão Conventional Commits
git add .
git commit -m "tipo: descrição concisa"  # Ex: feat: adiciona autenticação

git push origin feature/nome-da-feature
```

## 📝 Padrões do Projeto

### Código
- Siga o [Effective Go](https://golang.org/doc/effective_go.html)
- Use nomes descritivos para variáveis e funções
- Documente funções e tipos públicos
- Mantenha as funções pequenas e focadas
- Use `gofmt` ou `goimports` para formatação

### Testes
- Cobertura mínima de 80%
- Adicione testes para novas funcionalidades
- Execute todos os testes antes de enviar PRs

### Migrações de Banco de Dados
- Toda alteração no esquema requer migração
- Formato: `000N_descricao_da_migracao.up.sql`
- Inclua arquivo `.down.sql` para rollback
- Migrações devem ser idempotentes

#### Gerenciamento Manual de Migrações

Em alguns cenários, pode ser necessário gerenciar as migrações manualmente. Aqui está como fazer isso:

1. **Instalação do golang-migrate**
   ```bash
   # Linux (usando curl)
   curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
   sudo mv migrate /usr/local/bin/
   
   # macOS (usando Homebrew)
   brew install golang-migrate
   
   # Windows (usando Chocolatey)
   choco install migrate
   ```

2. **Comandos Básicos**
   ```bash
   # Aplicar todas as migrações pendentes
   migrate -database postgres://user:pass@localhost:5432/dbname -path ./sql/migrations up
   
   # Reverter a última migração
   migrate -database postgres://user:pass@localhost:5432/dbname -path ./sql/migrations down 1
   
   # Forçar a versão específica (útil em desenvolvimento)
   migrate -database postgres://user:pass@localhost:5432/dbname -path ./sql/migrations force VERSION
   
   # Verificar versão atual
   migrate -database postgres://user:pass@localhost:5432/dbname -path ./sql/migrations version
   ```

3. **Dicas de Uso**
   - Substitua `user`, `pass`, `localhost`, `5432` e `dbname` pelas suas credenciais
   - Use a flag `-verbose` para ver o que está acontecendo
   - Para ambientes de produção, use variáveis de ambiente para as credenciais
   
4. **Solução de Problemas**
   - Se encontrar erros de permissão, verifique as credenciais do banco de dados
   - Para problemas de caminho, use o caminho absoluto para o diretório de migrações
   - Use `migrate -help` para ver todas as opções disponíveis

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

1. Atualize sua branch com a `main`
2. Certifique-se que os testes estão passando
3. Atualize o CHANGELOG.md
4. Envie o PR com descrição clara
5. Referencie a issue relacionada

## 📄 Licença

Contribuições são licenciadas sob [MIT](LICENSE).
