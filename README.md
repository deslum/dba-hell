dba-hell
===============================================
Ты администратор баз данных? Эта головоломка для тебя

**Запуск:**

```bash
make build-all
make start
```

**Доступы:**

1. RabbitMQ 

- URL: localhost:15672
- Login: rabbitmq
- Password: rabbitmq

2. Postgres 

- URI "postgres://dba-test:dba-test@localhost/dba_test?sslmode=disable"

**Архитектура**

**Producer**: сервис, гененрирующий значения от 0 до max от uint64 и отправляет RabbitMQ 
 сообщения формата
```json
 {
 	"id": 10,                     // сгенерированное значение
 	"name": "TODO: Create name",  // имя сообщения, пока как заглушка
 	"number": 2,                  // номер потока числом
 	"body": "Thread number 2",    // номер потока который сгенерировал сообщение и отправил в очередь
 	"ts": 1583506044              // время в timestamp когда было сгененрировано это сообщение 
 }
```

**Writer** сервис, берет из очереди сообщения и вставляет пачками в Postgres.


**Задача**

1. Увеличить производительность вставок в БД, чтобы скорость работы Writer была выше, чем скорость работы Producer.
При этом изменять код сервисов не нужно. Конфигурировать только Postgres.
2. Перенести индексы и БД на разные диски
3. Настроить avtovacuum. 

Успехов!!!

Есть предложения по улучшению? Пиши:

Email: windowod@gmail.com\
Twitter: @Randomazer\
Telegram: @Randomazer


