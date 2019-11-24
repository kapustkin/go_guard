package tests

import (
	"log"
	"net"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/kapustkin/go_guard/pkg/integration-tests/config"
	"github.com/kapustkin/go_guard/pkg/integration-tests/tests"
)

func FeatureContext(s *godog.Suite) {
	// загрузка конфига
	conf := config.InitConfig()
	// ожидание доступности сервиса с таймайутом
	if !isServerOnline(conf.RestServer) {
		log.Fatalf("rest server offline")
	}

	// инициализация тестов
	test := tests.Init(conf)
	// выход из сценария, если он завершился с ошибкой
	s.AfterScenario(func(data interface{}, err error) {
		if err != nil {
			log.Fatalf("%v", err)
		}
	})

	tests.ListOfTests(s, test)
}

func isServerOnline(url string) bool {
	timeout := time.Duration(int64(1 * time.Second))

	for i := 0; i < 10; i++ {
		_, err := net.DialTimeout("tcp", url, timeout)
		if err == nil {
			// задержка на всякий случай
			time.Sleep(1 * time.Second)
			return true
		}

		time.Sleep(3 * time.Second)
	}

	return false
}
