# Projeto GraphQL em Go

Este é um projeto de exemplo utilizando GraphQL em Go com a biblioteca `gqlgen`.

## Requisitos

- Docker
- Docker Compose

## Como iniciar o projeto

1. Clone o repositório:

```sh
git clone https://github.com/Lucassamuel97/GraphQL.git
cd GraphQL
```
2. Inicie o projeto:
```sh
docker-compose up --build
```
Isso irá construir e iniciar o projeto utilizando Docker Compose.

3. Como usar:

Após iniciar o projeto, a API GraphQL estará disponível em `http://localhost:8080/` Você pode acessar um playground do GraphQL em `http://localhost:8080/playground` para testar consultas e mutações.

4. Exemplos de Mutations e query

```go
mutation createCategory{
  createCategory(input: {name: "Tecnologdia", description: "cursos de tecnologia"}){
    id
    name
    description
  }
}

mutation createCourse{
  createCourse(input: {name: "Curso teste", description: "the best", categoryId: "<Cole aqui o ID da Category criada>"}){
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

query queryCategoriesWithCourses{
  categories{
    id
    name
    courses{
      id
      name
      description
    }
  }
}

query queryCourses{
  courses {
    id
    name
    description
  }
}

query queryCoursesWithCategory{
  courses {
    id
    name
    description
    category {
     	id
        name
        description
    }
  }
}
```