# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Versionamento Semântico](https://semver.org/spec/v2.0.0.html).

## [0.3.0] - 2025-08-13

### Adicionado
- Suporte a gerenciamento de cursos (CRUD)
- Relacionamento entre cursos e categorias
- Migrações de banco de dados automáticas
- Suporte ao golang-migrate para gerenciamento de esquema
- Documentação atualizada para novas funcionalidades

### Corrigido
- Ajustes no tratamento de erros nas operações de banco de dados
- Correção na formatação de respostas GraphQL

## [0.2.0] - 2025-08-12

### Adicionado
- Configuração inicial do projeto GraphQL com gqlgen
- Configuração do banco de dados PostgreSQL
- Estrutura básica para gerenciamento de categorias
- Documentação inicial do projeto (README, CHANGELOG, CONTRIBUTING)
- Suporte a variáveis de ambiente com godotenv e envconfig
- Docker Compose para ambiente de desenvolvimento

### Alterado
- Refatoração da função main para melhor organização do código
- Melhorias no tratamento de erros e logs
- Atualização das dependências do projeto

### Corrigido
- Correção no carregamento de variáveis de ambiente
- Correção no fechamento da conexão com o banco de dados
- Ajustes para atender aos requisitos do linter

## [0.1.0] - 2025-08-10
### Adicionado
- Versão inicial do projeto
