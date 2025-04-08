# Этап сборки
FROM golang:1.21-alpine AS builder

# Установка необходимых зависимостей для сборки
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o user_service .

# Этап запуска
FROM alpine:latest

# Установка необходимых пакетов для запуска
RUN apk --no-cache add ca-certificates tzdata

# Создание непривилегированного пользователя
RUN adduser -D -g '' appuser

# Установка рабочей директории
WORKDIR /app

# Копирование бинарного файла из этапа сборки
COPY --from=builder /app/user_service .

# Установка прав на выполнение
RUN chown -R appuser:appuser /app

# Переключение на непривилегированного пользователя
USER appuser

# Открытие порта
EXPOSE 8082

# Запуск приложения
CMD ["./user_service"]