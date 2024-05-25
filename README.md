## Автор
 Даньшин Семён
 
## Описание
GraphQL-api для работы с данными о постах и комментариях.

## Стек технологий
- Golang
- Gqlgen
- Gorm
- Bearer авторизация и аутентификация через JWT

## Примеры запросов:
### - Создание поста
```graphql
mutation {
    createPost(input: {
        title: "как писать графкукл"
        content: "сначла пишешь потом плачешь"
        authorId: "f2c41b3a-74c3-4805-a979-758d8684ef7e"
    }) {
        id
        title
        content
        author {
            id
            name
        }
    }
}
```
Ответ:
```json
{
    "data": {
        "createPost": {
            "id": "1",
            "title": "как писать графкукл",
            "content": "сначла пишешь потом плачешь",
            "author": {
                "id": "f2c41b3a-74c3-4805-a979-758d8684ef7e",
                "name": "Семён"
            }
        }
    }
}
```

### - Получение всех постов
```graphql
query allPosts {
    posts (limit: 10, offset: 10) {
        id
        title
        author {
            id
            name
        }
        comments {
            id
            content
            author {
                id
                name
            }
        }
    }
}
```
Ответ:
```json
{
    "data": {
        "posts": [
            {
                "id": "1",
                "title": "как писать графкукл",
                "author": {
                    "id": "f2c41b3a-74c3-4805-a979-758d8684ef7e",
                    "name": "Семён"
                },
                "comments": [
                    {
                        "id": "1",
                        "content": "круто",
                        "author": {
                            "id": "e5943cb2-4bfe-4379-a93a-d24a46f7d244",
                            "name": "Антон"
                        }
                    }
                ]
            }
        ]
    }
}
```