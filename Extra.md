# Instalações

## GraphQL

https://graphql.org/learn/execution/

### Playground
```graphql
mutation createCategory{
  createCategory(input: {
    name: "Arquitetura 3",
    description: "Cursos de Arquitetura de Software 3"
  }){
    id
    name
    description
  }
}

mutation createCourse{
  createCourse(input: {
    name: "Git e Github 3",
    description: "Curso de git com github",
    categoryId: "c457ac63-3ea6-47e6-94ce-34851e2c0ddc"
  }){
    id
    name
    description
  }
}

query queryCategories{
  categories{
    id
    name
    description
  }
}

query queryCourses{
  courses{
    id
    name
    description
  }
}

query queryCategoriesWithCourses{
  categories{
    id
    name
    description
    courses{
      id
      name
    }
  }
}

query queryCoursesWithCategory{
  courses{
    id
    name
    category{
      name
    }
  }
}
```

### gqlgen

```bash
# Linux
sudo apt-get install -y gqlgen

# Windows
choco install -y gqlgen

# macOS
brew install gqlgen

# Verificar instalação
gqlgen --version
```

### gqlgen init

```bash
gqlgen init
```

### gqlgen generate

```bash
gqlgen generate
```

## gRPC

https://grpc.io/docs/languages/go/quickstart/

### Protocol buffer compiler

```bash
# Linux
sudo apt-get install -y protobuf-compiler

# Windows
choco install -y protobuf

# macOS
brew install protobuf

# Verificar instalação
protoc --version
```

### Go plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### protoc generate
```bash
protoc --go_out=internal/application/adapter/grpc/v1/pb --go-grpc_out=internal/application/adapter/grpc/v1/pb internal/application/adapter/grpc/v1/proto/course.proto
```
