# возьмем один из самых свежих образов Go
FROM golang:1.19.0-alpine3.16

# alpine образ беден на стандартные инструменты, добавим их отдельной командой
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# задаем директорию приложения в контейнере
WORKDIR /app

# скачиваем в контейнер зависимости приложения
COPY go.mod ./
RUN go mod download

# копируем файлы текущей директории в контейнер
COPY . .

# билдим приложение
RUN go build -o main .

# так как приложение читает 8890 порт, откроем его внешнему миру
EXPOSE 8890

# команда запуска приложения в контейнере
CMD ["./main"]