# Утилита для миграций
---
### Задание дано в DigitalSpirit при обучении
---

Для запуска необходимо поднять базу данных, где будут проводиться миграции.
Прописать необходимую конфигурацию в etc/config.yaml для подключения к базе, а также путь к миграциям.

Запуск:

```go
go run main.go -c etc/config.yaml -mode [up|down|reset|version]
```

up - "прогнать" все скрипты миграции с текущей до последней
down - "откатить" версию миграции на 1 назад
reset - "откатить" полностью БД с текущей версии до 0
version - вывести в stdout текущую версию миграции
