# TZ 

Http сервис - key-value хранилище. Умеющий сохранять/обновлять, удалять, выводить по id и весь список  

## Параметры запуска

Для того, чтобы запустить проект, необходимо выполнить команду в корне проекта ```go run ./cmd/main.go```

Сервер запускается на порту  ```:8080```

## Примеры

Пример для сохранения/обновления. Для сохранения/обновления user'a необходимо передать id в строке запроса. Если такой user есть,
то он обновит поля, если нет - создаст нового, с id, который был указан в строке запроса
```
PUT  /users/{id}
```

```
{
	"name": "NewName",
	"email": "NewEmail"
}
```
Пример для удаления.
Для удаления user'a необходимо передать id в строке запроса
```
DELETE /users/{id}
```

Пример для вывода всех имеющихся пользователей 
```
GET /users
```

Пример для вывода конкретного пользователя. Для этого необходимо передать id user'а в строке запроса
```
GET /users/{id}
```