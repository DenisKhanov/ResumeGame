# Resume-game

**Resume-game** — это интерактивное резюме в виде небольшой консольной игры

## Возможности сервиса


## Требования

Для запуска потребуется:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/manual/make.html)

## Установка и запуск


1. **Сгенерируйте ключи сервера для mTLS:**

    ```bash
    make server-keys
    ```

2. **Сгенерируйте ключи клиента для mTLS:**

    ```bash
    make client-keys
    ```

### Запуск сервера

Выполните команду из корня проекта:

```bash
docker compose up
```

### Запуск клиента

Выполните команды из корня проекта:

```bash
make build
make run
```

## Базовое использование

## Технологии

- **mTLS**: защищенное соединение между клиентом и сервером
- **PostgreSQL**: база данных с партицированием данных

## Безопасность


## Контакты

Если у вас есть вопросы или предложения, обратитесь к разработчику проекта через [GitHub](https://github.com/DenisKhanov/ResumeGame).

---