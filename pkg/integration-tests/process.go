package tests

import (
	"log"

	"github.com/DATA-DOG/godog"
	"github.com/kapustkin/go_guard/pkg/integration-tests/config"
	"github.com/kapustkin/go_guard/pkg/integration-tests/tests"
)

func FeatureContext(s *godog.Suite) {
	// загрузка конфига
	сonf := config.InitConfig()
	// инициализация тестов
	test := tests.Init(сonf)
	// выход из сценария, если он завершился с ошибкой
	s.AfterScenario(func(data interface{}, err error) {
		if err != nil {
			log.Fatalf("%v", err)
		}
	})

	tests.ExecPingTest(s, test)

	//tests.ExecCreateTest(s, test)
	//tests.ExecEditTest(s, test)
	//tests.ExecRemoveTest(s, test)
}
