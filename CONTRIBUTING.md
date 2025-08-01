# Guia de Contribuição

Obrigado por considerar contribuir para este projeto! Aqui estão algumas diretrizes para te ajudar a começar.

## 📋 Como Contribuir

1. **Encontre uma Issue**
   - Verifique as [issues abertas](https://github.com/seu-usuario/seu-projeto/issues) para encontrar algo que você gostaria de trabalhar
   - Se não encontrar uma issue que atenda ao que você deseja implementar, sinta-se à vontade para abrir uma nova

2. **Faça um Fork do Repositório**
   ```bash
   git clone https://github.com/seu-usuario/seu-projeto.git
   cd seu-projeto
   ```

3. **Crie uma Branch para sua Feature**
   ```bash
   git checkout -b feature/nova-funcionalidade
   ```

4. **Desenvolva sua Feature**
   - Siga o padrão de código do projeto
   - Adicione testes para suas alterações
   - Atualize a documentação conforme necessário

5. **Execute os Testes**
   ```bash
   go test ./...
   ```

6. **Verifique o Lint**
   ```bash
   golangci-lint run
   ```

7. **Faça o Commit das suas Alterações**
   ```bash
   git add .
   git commit -m "feat: adiciona nova funcionalidade"
   ```

8. **Envie as Alterações**
   ```bash
   git push origin feature/nova-funcionalidade
   ```

9. **Abra um Pull Request**
   - Vá até o repositório original
   - Clique em "New Pull Request"
   - Descreva suas alterações e referencie a issue relacionada

## 🏗️ Padrões de Código

- Siga o [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` ou `goimports` para formatar o código
- Documente funções e tipos públicos
- Escreva testes para novas funcionalidades

## 🧪 Testes

- Testes unitários devem cobrir pelo menos 80% do código
- Execute todos os testes antes de enviar um PR:
  ```bash
  go test -v ./...
  ```

## 📝 Padrão de Commits

Utilizamos o [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` Nova funcionalidade
- `fix:` Correção de bugs
- `docs:` Alterações na documentação
- `style:` Formatação, ponto e vírgula, etc. (sem alteração de código)
- `refactor:` Refatoração de código
- `test:` Adicionando testes
- `chore:` Atualização de tarefas, configurações, etc.

## 🔒 Segurança

Se você encontrar uma vulnerabilidade de segurança, por favor, não abra uma issue. Em vez disso, envie um e-mail para security@example.com.

## 📄 Licença

Ao contribuir, você concorda que suas contribuições serão licenciadas sob a licença [MIT](LICENSE).
