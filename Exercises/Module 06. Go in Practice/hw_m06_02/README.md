## Решите задачу:

Для работы с различными сервисами в интернете и API, первостепенными являются навыки забора и чтения данных в различных форматах. Улучшим эти навыки с помощью этой задачи:

Один из URL нашего GoMetr - ``/health`` по которому отдаются данные в json формате, следующего вида
```json
{
    "status": "pass",
    "service_id": "MBPadmincity101",
    "checks": {
        "ping_mysql": {
            "component_id": "mysql",
            "component_type": "db",
            "status": "pass"
            } 
        } 
}
```

Ваша задача написать клиент, который сможет вызвать этот сервер по адресу переданному во флаг ```--url``` и выведет значения полей ```status```, ```service_id``` и ```checks.ping_mysql.status``` в виде **“Overall status is %status%, with service_id %значение поля service_id% mysql component is %начение поля checks.ping_mysql.status%”**, либо **“No data”** если до сервера не удалось достучаться.