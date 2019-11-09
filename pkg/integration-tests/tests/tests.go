package tests

import (
	"github.com/DATA-DOG/godog"
)

func ListOfTests(s *godog.Suite, test *NotifyTest) {
	s.Step(`^посылаю "([^"]*)" запрос к "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^ожидаю что код ответа будет (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^тело ответа будет равно "([^"]*)"$`, test.theResponseShouldMatchText)
	s.Step(`^посылаю "([^"]*)" запрос к "([^"]*)" c "([^"]*)" содержимым:$`, test.theSendRequestToWithData)
	s.Step(`^посылаю "([^"]*)" запрос к "([^"]*)" в количестве (\d+) раз c "([^"]*)" содержимым:$`, test.theSendManyRequestsToWithData)
	s.Step(`^ответ тело ответа будет с содержимым:$`, test.theResponseDocStringShouldMatchText)
}
