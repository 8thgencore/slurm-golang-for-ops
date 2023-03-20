# Service

Сервис парсит метрики у ресурса gometr и размещает в своей базе данных

## Запуск

Запустить сбор метрик
```bash
go run cmd/main.go
```

Получить собранные метрики
```bash
http://ip:8350/metrics/list
```


Список фильтров:
```toml
name=go_gc_duration_seconds_count
time_from=2023-03-14T19:57:40.106073Z
time_to=2023-03-15T19:57:40.106073Z
limit=5
offset=10
```

Например
```
http://ip:8350/metrics/list?name=go_gc_duration_seconds_count&time_from=2023-03-14T19:57:40.106073Z&limit=5&offset=10
```
