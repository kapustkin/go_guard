package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal/checker"
	logger "github.com/sirupsen/logrus"
)

// Check all events for user
func (handler *MainHandler) RequestChecker(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error read request body", http.StatusForbidden)
		return
	}
	var r struct {
		Login    string
		Password string
		IP       string
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error parsing body", http.StatusForbidden)
		return
	}
	logger.Infof("process request %v", r)

	checkResult, err := listsChecks(handler, r.IP, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if checkResult {
		return
	}

	err = mainChecks(handler, r.Login, r.Password, r.IP, res)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

//mainChecks in buckets
func mainChecks(handler *MainHandler, login, pass, addr string, res io.Writer) error {
	// get parameters from db
	params, err := handler.db.GetParametrs()
	if err != nil {
		logger.Errorf(err.Error())
		return fmt.Errorf("error load parameters %v", err)
	}
	//check login
	loginRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("l_%s", login), params.K)
	if err != nil {
		logger.Errorf(err.Error())
		return fmt.Errorf("error in process bucket %v", err)
	}
	//check password
	passwordRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("p_%s", pass), params.M)
	if err != nil {
		logger.Errorf(err.Error())
		return fmt.Errorf("error in process bucket %v", err)
	}
	//check ip
	ipRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("i_%s", addr), params.N)
	if err != nil {
		logger.Errorf(err.Error())
		return fmt.Errorf("error in process bucket %v", err)
	}
	//nolint:
	res.Write([]byte(fmt.Sprintf("ok=%v", loginRes && passwordRes && ipRes)))
	logger.Infof("login:%v=%v, pass:%v=%v, ip:%v=%v, result=>%v",
		login, loginRes, pass, passwordRes, addr, ipRes,
		loginRes && passwordRes && ipRes)
	return nil
}

//listsChecks in black/white lists
func listsChecks(handler *MainHandler, ip string, res io.Writer) (bool, error) {
	//black-white list ip check
	lists, err := handler.db.GetAddressList()
	if err != nil {
		logger.Errorf(err.Error())
		return false, fmt.Errorf("error load white/black lists %v", err)
	}
	for _, item := range *lists {
		result, err := checker.IsAddressInNewtork(item.Network, ip)
		if err != nil {
			logger.Errorf(err.Error())
		}
		if result {
			logger.Infof("ip address %v contains in %v", ip, item.Network)
			//nolint:
			res.Write([]byte(fmt.Sprintf("ok=%v", item.IsWhite)))
			logger.Infof("process result %v=%v", ip, item.IsWhite)
			return true, nil
		}
	}
	return false, nil
}
