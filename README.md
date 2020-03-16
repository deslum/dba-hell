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

![](https://gitlab.com/deslum/dba-hell/-/blob/master/arch.png)

**Producer**: сервис, гененрирующий значения от 0 до max от uint64 и отправляет RabbitMQ 
 сообщения формата
```json
 {
 	"id": 10,                     
 	"name": "TODO: Create name",   
 	"number": 2,                   
 	"body": "Thread number 2",     
 	"ts": 1583506044              
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



