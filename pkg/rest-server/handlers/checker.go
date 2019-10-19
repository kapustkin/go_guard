package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal/checker"
	logger "github.com/sirupsen/logrus"
)

type request struct {
	Login    string
	Password string
	IP       string
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
	// get parameters from db
	params, err := handler.db.GetParametrs()
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error load parameters", http.StatusForbidden)
		return
	}
	//black-white list ip check

	//usual checks
	loginRes, err := checker.ProcessBucket(handler.store, r.Login, params.K)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error in processBucket", http.StatusInternalServerError)
		return
	}
	logger.Infof("process result %v=%v", r.Login, loginRes)
	passwordRes, err := checker.ProcessBucket(handler.store, r.Password, params.M)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error in processBucket", http.StatusInternalServerError)
		return
	}
	logger.Infof("process result %v=%v", r.Password, passwordRes)
	ipRes, err := checker.ProcessBucket(handler.store, r.IP, params.N)
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
