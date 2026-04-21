# Shikimori Ongoings API

Небольшой сервис на Go, который парсит календарь онгоингов с [Shikimori](https://shikimori.one/ongoings) и отдаёт список аниме, выходящих через N дней от текущей даты.

## Стек

- Go 1.22+
- Gin (HTTP)
- Goquery (парсинг HTML)

## Запуск

```bash
git clone https://github.com/Reteger/shiki.git
cd shiki

go mod tidy
go run cmd/server/main.go
```

По умолчанию сервер поднимается на `http://localhost:8080`.

## Эндпоинт

### GET `/api/ongoings/:days`

Возвращает онгоинги, которые выходят через `days` дней от сегодняшнего дня.

- `days` — целое число от 1 до 7.

#### Пример

```bash
curl "http://localhost:8080/api/ongoings/1"
```

Ответ:

```json
{
  "day": "Вторник",
  "days_ahead": 1,
  "titles": [
    {
      "original_title": "Fushigi Dagashiya: Zenitendou",
      "russian_title": "Таинственный магазин сладостей «Дзэнитэндо»",
      "link": "https://shikimori.one/animes/42295-fushigi-dagashiya-zenitendou"
    }
  ]
}
```
##  ❗ Важно,  Ограничения 

- Используется парсинг HTML Shikimori, при изменении разметки сайта код может потребовать правок.
- Для доступа к `shikimori.one` в РФ  требуется VPN.
