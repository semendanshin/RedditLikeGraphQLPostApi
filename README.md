## Автор
 Даньшин Семён
 
## Описание проекта
GraphQL-api для работы с данными о постах и комментариях.

## Детали реализации
Реализовано хранение данных в оперативной памяти и в базе данных PostgreSQL 
(переключается в файле конфигурации ./config/example.yaml). 
Для работы с данными используется ORM Gorm.
Репозитории реализованы композицией с абстрактным репозиторием, что позволяет избежать дублирования кода, сохраняя при этом гибкость и расширяемость.



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
            "id": "af9604be-a817-49ff-8cc4-0922cc0f63aa",
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
                "id": "af9604be-a817-49ff-8cc4-0922cc0f63aa",
                "title": "как писать графкукл",
                "author": {
                    "id": "f2c41b3a-74c3-4805-a979-758d8684ef7e",
                    "name": "Семён"
                },
                "comments": [
                    {
                        "id": "9aa6d910-ff6e-4081-bca7-2fcd0f73a661",
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