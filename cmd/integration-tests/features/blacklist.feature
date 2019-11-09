# file: features/blacklist.feature

Feature: Проверка адреса, который находится в черном списке
        Когда придет запрос на проверку адреса http://{REST_SERVER}/checker
        Если адрес находится в черном спике
        Тогда всегда возвращать ok=false

        Scenario: Сервис Анти-брутфорс доступен
                When посылаю "GET" запрос к "http://{REST_SERVER}/ping"
                Then ожидаю что код ответа будет 200
                And тело ответа будет равно "OK"

        Scenario: Добавляем узел в черный список
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/lists/add" c "application/json" содержимым:
		"""
                        {
                        "Network": "127.0.14.1/30",
                        "IsWhite": false
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
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """

        Scenario: Делаем запросы для проверки работоспособности blacklist
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test2",
                        "password": "4321",
                        "ip": "127.0.14.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """

        Scenario: Удаляем адрес из черного списка
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/lists/remove" c "application/json" содержимым:
		"""
                {
                "Network": "127.0.14.1/30",
                "IsWhite": false
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
                        "login": "test2",
                        "password": "4321",
                        "ip": "127.0.14.2"
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
                        "login": "test2",
                        "password": "4321",
                        "ip": "127.0.14.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """