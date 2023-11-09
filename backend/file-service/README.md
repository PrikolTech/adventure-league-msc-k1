# Файловый сервер
ruby 3.2 + sinatra

### Эндпоинты

#### Отправка и сохранение файла
```
 /api/files?name={string} (POST)

Response:
{
  url: string
}
```

#### Получение файла
```
> /api/files/{url} (GET) 

Response:
file
```
