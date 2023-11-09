# Микросервис Курсов

## Запуск
Требуется докер
```
> docker-compose build 
> docker-compose up 
```

*Микросервис будет поднят по адресу localhost*:**3003**

Подготовка базы данных (с предварительным заполнением)
```
> docker-compose exec api bundle exec rake db:drop db:create db:migrate db:seed
```

## Эндпоинты

### Схемы

#### Курс
```
Course: {
    "id": uuid,
    "name": string,
    "description": string,
    "advantages": string,
    "price": float || null,
    "tg_link": string || null,
    "prefix": string,
    "created_at": datetime,
    "updated_at": datetime,
    "period": {
      "starts_at": date,
      "ends_at": date
    },
    "course_type": {
      "name": string
    },
    "education_form": {
      "name": string
    }
  } 
```

#### Лекция
```
Lecture: {
  "id": uuid,
  "name": string,
  "description": string,
  "starts_at": datetime,
  "course_id": uuid,
  "is_hidden": boolean,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Группа
```
Group: {
  "id": uuid,
  "name": string,
  "course_id": uuid,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Пользователь-Группа
```
UserGroup: {
  "group_id": uuid,
  "user_id": uuid,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Контент (файлы)

```
Content: {
  "id": uuid,
  "url": string,
  "lecture_id": "11111111-1111-1111-1111-111111111111",
  "created_at": "2023-11-09T13:07:10.800Z",
  "updated_at": "2023-11-09T13:07:10.800Z",
  "content_type": {
    "id": "0279e5bd-266f-40d2-9abc-0c39a31d8801",
    "name": "TEST FILE"
  }
}
```

### Курсы

```
/api/courses?user_id={user_id} (GET)

Response:
[
  Course,
  Course,
  ...
]
```

```
/api/courses (POST) -> Create
/api/courses/{course_id} (PUT) -> Update

Request: 
{
  Course
}

Response:
{
  Course {
    "lectures": [ 
      Lecture,
      Lecture
    ]
  }
}

```

Добавление пользователя к курсу
```
/api/courses/{course_id}/add_user?user_id={user_id}

Response:
{
  Group
}
```

```
/api/courses/{course_id} (DELETE)
```

### Лекции

```
/api/courses/{course_id}/lectures?user_id={user_id}(GET)

Response:
[
  Lecture,
  Lecture,
  ...
]
```

```
/api/courses/{course_id}/lectures (POST) -> Create
/api/courses/{course_id}/lectures/{lecture_id} (PUT) -> Update

Request: 
{
  Lecture
}

Response:
{
  Lecture
}

```

```
/api/courses/{course_id}/lectures/{lecture_id} (DELETE)
```

### Контент (файлы)

```
/api/courses/{course_id}/lectures/{lecture_id}/contents (GET)

Response:
[
  Content,
  Content,
  ...
]
```

```
/api/courses/{course_id}/lectures/{lecture_id}/contents (POST)

Request:
file

Response:
{
  Content
}
```

```
/api/courses/{course_id}/lectures/{lecture_id}/contents/{content_id} (DELETE)
```



### Группы

```
/api/groups/{course_id}?user_id (GET)

Response:
[
  Group,
  Group,
  ...
]
```

```
/api/groups/{course_id} (POST)
/api/groups/{course_id} (PUT)

Request:
{
  Group
}

Response:
{
  Group {
    "user_groups": [
      UserGroup,
      UserGroup,
      ...
    ]
  }
}
```

```
/api/groups/{course_id} (DELETE)
```