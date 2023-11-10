# Микросервис Тестирований и Заданий

### Запуск
Требуется докер
```
> docker-compose build 
> docker-compose up 
```
*Микросервис будет поднят по адресу localhost*:**3004**

### Схемы

#### Работа

```
Job: {
    "id": uuid,
    "lecture_id": uuid,
    "name": string,
    "description": string,
    "deadline": datetime,
    "job_type_id": integer,
    "created_at": datetime,
    "updated_at": datetime
  }
```

#### Тест

```
Test: {
  "id": uuid,
  "job_id": uuid,
  "name": string,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Вопрос

```
Question: {
  "id": uuid,
  "test_id": uuid,
  "body": string, # Сам текст вопроса
  "is_multiple_answer": boolean,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Ответ

```
Aswer: {
  "id": uuid,
  "question_id": uuid,
  "is_correct": boolean,
  "body": string,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Решение теста

```
TestSolution: {
  "id": uuid,
  "test_id": uuid,
  "user_id": uuid,
  "created_at": datetime,
  "updated_at": datetime,
  "test_result": {
    "id": uuid,
    "test_solution_id": uuid,
    "score": integer, # Количество правильных ответов
    "created_at": datetime,
    "updated_at": datetime
  },
  "solution_answers": [НЕ ОБРАБАТЫВАТЬ]
}
```

#### Домашнее задание

```
Homework: {
  "id": uuid,
  "job_id": uuid,
  "name": string,
  "description": string,
  "created_at": datetime,
  "updated_at": datetime
}
```

#### Решение домашнего задания

```
HomeworkSolution: {
  "id": uuid,
  "homework_id": uuid,
  "user_id": uuid,
  "file_url": string,
  "created_at": datetime,
  "updated_at": datetime
}
```


### Работы

```
  api/jobs?lecture_id (GET)

  Response:
  [
    Job,
    Job,
    ...
  ]
```

```
  api/jobs (POST)

  Request:
  {
    Job
  }

  Response:
  {
    Job
  }
```

```
  api/jobs/{job_id} (GET)

  Response:
  {
    Job,
    "tests": [
      Test,
      Test
    ]
  }
```

### Тесты

```
  api/jobs/{job_id}/tests (GET)

  Response:
  [
    Test,
    Test,
    ...
  ]
```
```
  api/jobs/{job_id}/tests (POST)

  Request:
  {
    Test
  }

  Response:
  {
    Test
  }
```

```
  api/jobs/{job_id}/tests/{test_id} (GET)

  Response:
  {
    Test,
    "questions": [
      Question,
      Question
    ]
  }
```

### Решение теста

```
  api/jobs/{job_id}/tests/{test_id}/test_solutions?user_id={user_id} (GET)

  Response:
  [
    TestSolution,
    TestSolution,
    ...
  ]
```
```
  api/jobs/{job_id}/tests/{test_id}/test_solutions (POST)

  Request:
  {
    "user_id": uuid,
    answers: [
      {
        "answer_id": uuid
      },
      {
        "answer_id": uuid
      }
    ]
  }

  Response:
  {
    TestSolution
  }
```

```
  api/jobs/{job_id}/tests/{test_id}/test_solutions (GET)

  Response:
  {
    Test,
    "questions": [
      Question,
      Question
    ]
  }
```

### Вопросы

```
  api/jobs/{job_id}/tests/{test_id}/questions (GET)

  Response:
  [
    Questions,
    Questions,
    ...
  ]
```
```
  api/jobs/{job_id}/tests/{test_id}/questions (GET)

  Request:
  {
    Question
  }

  Response:
  {
    Question
  }
```
```
  api/jobs/{job_id}/tests/{test_id}/questions/{question_id} (GET)

  Response:
  [
    Question,
    "answers": [
      Answer,
      Answer,
      ...
    ]
  ]
```

### Домашняя работа

```
api/jobs/{job_id}/homeworks (GET, POST)
```

### Решение домашней работы

```
api/jobs/{job_id}/homeworks/{homework_id}/homework_solutions (GET, POST)
```

### Отправить оценку

```
api/jobs/{job_id}/homeworks/{homework_id}/homework_solutions/{solutions_id}/result (POST)

Request:
{
  "score": integer # From 2 to 5
}
```
