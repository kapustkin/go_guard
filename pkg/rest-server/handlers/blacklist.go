package handlers

import (
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func (e *MainHandler) AddToBlackList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (e *MainHandler) RemoveFromBlackList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
