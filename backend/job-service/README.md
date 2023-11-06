# Микросервис Тестирований и Заданий

### Запуск
Требуется докер
```
> docker-compose build 
> docker-compose up 
```

### Подготовка Базы Данных
```
> docker-compose exec api bundle exec rake db:drop db:create db:migrate db:seed
```