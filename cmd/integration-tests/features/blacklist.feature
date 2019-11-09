# file: features/blacklist.feature

Feature: Проверка адреса, который находится в белом списке
        Когда придет запрос на проверку адреса http://{REST_SERVER}/checker
        Если адрес находится в белом спике
        Тогда всегда возвращать ok=true

        Scenario: Сервис Анти-брутфорс доступен
                When посылаю "GET" запрос к "http://{REST_SERVER}/ping"
                Then ожидаю что код ответа будет 200
                And тело ответа будет равно "OK"