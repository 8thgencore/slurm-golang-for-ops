Для создания таблицы в базе данных для хранения таких структур можно использовать следующий SQL-запрос:

```sql
CREATE TABLE metrics (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL
);
```

Это создаст таблицу с четырьмя полями: `id` - уникальный идентификатор метрики, `timestamp` - время забора метрики, `name` - имя метрики и `value` - значение метрики. 

Для ускорения поиска по имени или времени можно добавить индексы на эти поля:

```sql
CREATE INDEX idx_name ON metrics (name);
CREATE INDEX idx_timestamp ON metrics (timestamp);
```