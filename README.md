# README

## Описание проекта

Этот проект представляет собой сервис, написанный на языке Go, который предоставляет пользователю API для вычисления математических выражений

## Установка

Для начала, убедитесь, что у вас установлен Go. Если его нет, вы можете скачать и установить Go с [официального сайта](https://golang.org/dl/).

Затем, клонируйте репозиторий:

```bash
git clone https://github.com/ваш_username/ваш_репозиторий.git
cd ваш_репозиторий
```

Установите зависимости:

```bash
go mod tidy
```

## Использование

### Запуск сервиса

Для запуска сервиса используйте следующую команду:

```bash
go run cmd/main.go
```

Сервис будет запущен на `http://localhost:8080` по умолчанию

Для смены порта необходимо указать его в окружении перед запуском
   Для Windows command prompt:
      ```
      set PORT=3000
      go run main.go
      ```
   Для PowerShell:
      ```
      $env:PORT=3000
      go run main.go
      ```
   Если вы используете Docker, вы можете установить переменную окружения в Dockerfile:
      ```
      ENV PORT=3000
      ```

   Либо передать переменные окружения при запуске контейнера:
      ```bash
      docker run -e PORT=3000 your-image-name
      ```

### API

Сервис предоставляет один **endpoint**:

- `POST /api/v1/calculate`

#### Запрос

Данный endpoint принимает POST-запрос с JSON-телом:

```json
{
    "expression": "выражение, которое ввёл пользователь"
}
```

Где `expression` — это строка, содержащая математическое выражение, которое нужно вычислить.

#### Ответы

В зависимости от результата вычисления можно получить несколько вариантов ответов:

1. **Успех**

   Если выражение вычислено успешно, вы получите ответ с кодом 200:

   ```json
   {
       "result": "результат выражения"
   }
   ```

2. **Некорректное выражение**

   Если входные данные не соответствуют требованиям (например, содержат недопустимые символы), вы получите ответ с кодом 422:

   ```json
   {
       "error": "Expression is not valid"
   }
   ```

3. **Внутренняя ошибка сервера**

   В случае возникновения любой другой ошибки, вы получите ответ с кодом 500:

   ```json
   {
       "error": "Internal server error"
   }
   ```

### Пример запроса

Пример запроса с использованием `curl`:

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "3 + 5"}'
```

### Пример успешного ответа

```json
{
    "result": "8"
}
```

### Пример некорректного выражения

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "3 + a"}'
```

Ответ:

```json
{
    "error": "Expression is not valid"
}
```

## Тестирование

Для тестирования проекта можно использовать встроенные пакеты `testing`. Запустите команду:

```bash
go test ./...
```