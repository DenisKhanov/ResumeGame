# Stage 1: Builder
FROM golang:1.22.2-alpine AS builder

# Устанавливаем `goose` внутри контейнера
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Создаем и переходим в директорию приложения
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Копируем server tlsconfig.
COPY pkg/tlsconfig/cert/server ./

# Загружаем все зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию контейнера
COPY ./ ./

# Копируем файл .env
COPY server.env ./

# Сборка бинарного файла
RUN CGO_ENABLED=0 GOOS=linux go build -a -o resumegamesrv ./cmd/server/resume_server.go

# Stage 2: Runner
FROM alpine:3.19

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранный бинарник и файл .env из предыдущего этапа
COPY --from=builder /app/resumegamesrv .
COPY --from=builder /app/server.env ./
COPY --from=builder /app/pkg/tlsconfig/cert/server /app/pkg/tlsconfig/cert/server/

# Проверка наличия сертификатов
RUN ls -la /app/pkg/tlsconfig/cert/server/

# Запускаем бинарник
CMD ["./resumegamesrv"]