package tests

import (
	"github.com/DATA-DOG/godog"
)

func ExecPingTest(s *godog.Suite, test *NotifyTest) {
	s.Step(`^посылаю "([^"]*)" запрос к "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^ожидаю что код ответа будет (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^тело ответа будет равно "([^"]*)"$`, test.theResponseShouldMatchText)
}
