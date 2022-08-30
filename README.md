ТЗ:
Реализовать сервис приема логов. Основной упор в реализации должен быть на большое кол-во rps. Можно использовать любые библиотеки и решения. Ограничений по памяти нет.
Сервис принимает HTTP POST запрос на url /log
Лог приходит в формате json:

{
    user_uuid: UUID,
    timestamp: unixtimestamp,
    events: [{
        url: request url,
        dataRequest: request payload,
        dataResponse: response payload 
    }]
}


Лог необходимо принять, добавить в него уникальный uuid записи, IP адрес и записать в любую БД (на ваш выбор: clickhouse/mysql/mongodb/redis)
