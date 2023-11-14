# Лига приключений - Корпоративный университет

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

В проекте реализован ролевой доступ. Для тестирования каждой роли можно воспользоваться соответствующими учетными данными вида:
```
Логин: <роль>@mail.ru
Пароль: aA$123123
```

Список ролей:
- admin (администратор)
- employee (сотрудник)
- tutor (куратор)
- teacher (преподаватель)
- student (студент)
- enrollee (абитуриент)

Например, учетные данные сотрудника выглядят так:
```
Логин: employee@mail.ru
Пароль: aA$123123
```

## Доступный функционал
`employee (сотрудник)` может:
- Просматривать заявки по курсам
- Изменять статус заявки ("принята", "одобрена", "в резерве")

Изменение статуса заявки на `"принята"` ведет к автоматическому созданию аккаунта пользователя (если у него еще нет аккаунта) со случайно сгенерированным паролем. Учетные данные отправляются на указанную в заявке почту.

Изменение статуса заявки на `"одобрена"` ведет к автоматическому добавлению пользователя в выбранную специальным алгоритмом группу, привязанную к указанному в заявке курсу.

`tutor (куратор)` может:
- Создать курс
- Создать лекцию по курсу

`teacher (преподаватель)` может:
- Редактировать информацию в лекции по курсу
- Добавлять задания к лекции
- Прикреплять файлы к лекции

`student (студент)` может:
- Добавлять решение задания (прикреплять файл)
- Просматривать лекцию
- Проходить тесты
- Просматривать свои оценки
- Просматривать свои заявки на курсы
- Отправлять заявки на другие курсы

`enrollee (абитуриент)` может:
- Просматривать свои заявки на курсы
- Отправлять заявки на другие курсы

`Пользователь без аккаунта (анонимный пользователь)` может: 
- Отправлять заявки на курсы

Также к лекциям и заданиям можно оставлять комментарии.
