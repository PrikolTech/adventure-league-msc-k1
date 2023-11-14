# PrikolTech

## Установка и запуск
Для корректной работы проекта необходимо запустить обе части.

### Серверная часть
```
docker compose -f backend/docker-compose.yaml up --build 
```

### Клиентская часть
```
docker compose -f frontend/docker-compose.yaml up --build
```

## Взаимодействие и тестирование
Взаимодействие осуществляется через клиентскую часть по адресу: [localhost:5173](http://localhost:5173).

Тестирование функционала, связанного с почтовым сервисом, осуществляется в [Debugmail](https://app.debugmail.io/app/teams/linguine/projects/adventure-league) под учетными данными:
```
Логин: yph98384@omeie.com
Пароль: n7PepEHe6B
```