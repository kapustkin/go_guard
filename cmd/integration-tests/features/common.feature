# file: features/common.feature

Feature: Проверка основного алгоритма
        Когда придет запрос на проверку адреса http://{REST_SERVER}/checker
        Если адрес не превышены лимита KMN
        Тогда возвращать ok=true
        Иначе возвращать ok=false

        Scenario: Сервис Анти-брутфорс доступен
                When посылаю "GET" запрос к "http://{REST_SERVER}/ping"
                Then ожидаю что код ответа будет 200
                And тело ответа будет равно "OK"
        
        Scenario: Устанавливаем лимиты
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "n": 5,
                        "m": 100,
                        "k": 100
                        }
		"""
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """
        
        Scenario: Проверка N (логин). Делаем 5 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "testN",
                        "password": "password",
                        "ip": "127.0.10.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка N (логин). Делаем 5 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "testN",
                        "password": "password",
                        "ip": "127.0.10.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """

        Scenario: Устанавливаем лимиты
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "n": 100,
                        "m": 5,
                        "k": 100
                        }
		"""
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """

        Scenario: Проверка M (пароль). Делаем 5 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test",
                        "password": "passwordN",
                        "ip": "127.0.10.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка M (пароль). Делаем 5 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test",
                        "password": "passwordN",
                        "ip": "127.0.10.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """

        Scenario: Устанавливаем лимиты
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "n": 100,
                        "m": 100,
                        "k": 5
                        }
		"""
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """        
        
        Scenario: Проверка K (адрес). Делаем 5 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test2",
                        "password": "password3",
                        "ip": "127.0.10.5"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка K (адрес). Делаем 5 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "test2",
                        "password": "password3",
                        "ip": "127.0.10.5"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """