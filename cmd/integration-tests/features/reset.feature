# file: features/reset.feature

Feature: Проверка функциональности сброса данных бакета
        Когда придет запрос на сброс бакета по адресу http://{REST_SERVER}/admin/reset
        Если адрес или логин заблокирован
        Тогда разблокировать адрес и логин

        Scenario: Сервис Анти-брутфорс доступен
                When посылаю "GET" запрос к "http://{REST_SERVER}/ping"
                Then ожидаю что код ответа будет 200
                And тело ответа будет равно "OK"

        Scenario: Устанавливаем лимиты для логина
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "n": 5,
                        "m": 100,
                        "k": 100
                        }
		"""

        Scenario: Проверка N (логин). Делаем 5 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestN",
                        "password": "resetpassword",
                        "ip": "127.0.20.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка N (логин). Делаем 10 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestN",
                        "password": "resetpassword",
                        "ip": "127.0.20.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """

        Scenario: Сбрасываем данные бакета
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/reset" c "application/json" содержимым:
		"""
                        {
                        "login": "resettestN",
                        "ip": "127.0.20.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """

        Scenario: После сброса бакета делаем 5 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 5 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestN",
                        "password": "resetpassword",
                        "ip": "127.0.20.2"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """


        Scenario: Устанавливаем лимиты для адреса
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/params" c "application/json" содержимым:
		"""
                        {
                        "n": 100,
                        "m": 100,
                        "k": 10
                        }
		"""

        Scenario: Проверка К (адрес). Делаем 10 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestK",
                        "password": "resetpassword",
                        "ip": "127.0.20.3"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка К (адрес). Делаем 10 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestK",
                        "password": "resetpassword",
                        "ip": "127.0.20.3"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """

        Scenario: Сбрасываем данные бакета
                When посылаю "POST" запрос к "http://{REST_SERVER}/admin/reset" c "application/json" содержимым:
		"""
                        {
                        "login": "resettestK",
                        "ip": "127.0.20.3"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true}
                """

        Scenario: Проверка К (адрес). Делаем 10 запросов, ожидаем ответ ok=true
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestK",
                        "password": "resetpassword",
                        "ip": "127.0.20.3"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=true"}
                """

        Scenario: Проверка К (адрес). Делаем 10 запросов, ожидаем ответ ok=false
                When посылаю "POST" запрос к "http://{REST_SERVER}/checker" в количестве 10 раз c "application/json" содержимым:
		"""
                        {
                        "login": "resettestK",
                        "password": "resetpassword",
                        "ip": "127.0.20.3"
                        }
		"""        
                Then ожидаю что код ответа будет 200
                And ответ тело ответа будет с содержимым:
                """
                {"success":true,"data":"ok=false"}
                """