# file: features/common.feature

Feature: Проверка адреса, который находится в белом списке
        Когда придет запрос на проверку адреса http://{REST_SERVER}/checker
        Если адрес находится в белом спике
        Тогда всегда возвращать ok=true

        Scenario: Сервис Анти-брутфорс доступен
                When посылаю "GET" запрос к "http://{REST_SERVER}/ping"
                Then ожидаю что код ответа будет 200
                And тело ответа будет равно "OK"

        Scenario: Добавляем узел в белый список
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/lists/add" c "application/json" содержимым:
		"""
                        {
                        "Network": "127.0.13.1/30",
                        "IsWhite": true
                        }
		"""
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """
        
        Scenario: Устанавливаем лимиты
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "k": 2,
                        "m": 2,
                        "n": 2
                        }
		"""
                Then ожидаю что код ответа будет 200

        Scenario: Делаем запросы для проверки работоспособности whitelist
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 20 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test1",
                        "password": "12314",
                        "ip": "127.0.13.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Удаляем адрес из белого списка
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/lists/remove" c "application/json" содержимым:
		"""
                {
                "Network": "127.0.13.1/30",
                "IsWhite": true
                }
		"""
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """


        Scenario: Делаем 2 запроса для проверки, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 2 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test1",
                        "password": "12314",
                        "ip": "127.0.13.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """
        
        Scenario: Делаем запросы для проверки, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test1",
                        "password": "12314",
                        "ip": "127.0.13.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """