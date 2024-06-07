# Используем официальный образ golang для сборки приложения
FROM golang:1.17-alpine AS builder

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем модули Go и список зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение
RUN go build -o main .

# Фаза production - создаем минимальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/main .

# Запускаем приложение
CMD ["./main"]
