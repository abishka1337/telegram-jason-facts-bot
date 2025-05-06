Telegram Fact Bot (на Go)
Простой Telegram-бот на Go, который отвечает случайным фактом из PostgreSQL при команде /fact.

Возможности
Получение обновлений от Telegram API

Ответ пользователю случайным фактом из БД

Поддержка команды /fact

Технологии
Go

PostgreSQL

Telegram Bot API

HTTP-запросы, JSON

Как запустить
Установи PostgreSQL и создай таблицу:

sql
Копировать
Редактировать
CREATE TABLE facts (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL
);
Добавь туда хотя бы пару фактов:

sql
Копировать
Редактировать
INSERT INTO facts (text) VALUES ('Слон — единственное животное, которое не может прыгать.');
INSERT INTO facts (text) VALUES ('Мед у пчёл никогда не портится.');
Проверь, что строка подключения к БД соответствует твоей системе:

go
Копировать
Редактировать
const dbConnStr = "user=postgres password=123 dbname=telegram_bot1 sslmode=disable"
Замени botToken на свой токен от BotFather в Telegram:

go
Копировать
Редактировать
const botToken = "your_token_here"
Запусти проект:

bash
Копировать
Редактировать
go run main.go
Как использовать
В Telegram-боте отправь команду /fact, и бот пришлёт случайный факт из базы данных.
