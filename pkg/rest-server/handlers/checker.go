package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	storage "github.com/kapustkin/go_guard/pkg/rest-server/dal"
	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal"
	logger "github.com/sirupsen/logrus"
)

type MainHandler struct {
	db      storage.Storage
	k, m, n int
}

type request struct {
	Login    string
	Password string
	IP       string
}

// Init main handler
func Init(st *storage.Storage) *MainHandler {
	return &MainHandler{db: *st, k: 3, m: 3, n: 3}
}

// Check all events for user
func (handler *MainHandler) RequestChecker(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error read request body", http.StatusForbidden)
		return
	}
	var r request
	err = json.Unmarshal(body, &r)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error parsing body", http.StatusForbidden)
		return
	}
	logger.Infof("process request %v", r)
	//black-white list ip check

	//usual checks
	loginRes, err := internal.ProcessBucket(handler.db, r.Login, handler.k)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error in processBucket", http.StatusInternalServerError)
		return
	}
	logger.Infof("process result %v=%v", r.Login, loginRes)
	passwordRes, err := internal.ProcessBucket(handler.db, r.Password, handler.m)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error in processBucket", http.StatusInternalServerError)
		return
	}
	logger.Infof("process result %v=%v", r.Password, passwordRes)
	ipRes, err := internal.ProcessBucket(handler.db, r.IP, handler.n)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error in processBucket", http.StatusInternalServerError)
		return
	}
	logger.Infof("process result %v=%v", r.IP, ipRes)
	//nolint:
	res.Write([]byte(fmt.Sprintf("ok=%v", loginRes && passwordRes && ipRes)))
	logger.Infof("response ok=%v", loginRes && passwordRes && ipRes)
}

func (handler *MainHandler) ResetBucket(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
