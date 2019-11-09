package handlers

import (
	"fmt"
	"net/http"

	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal"
	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal/checker"
	logger "github.com/sirupsen/logrus"
)

// Check all events for user
func (handler *MainHandler) RequestChecker(res http.ResponseWriter, req *http.Request) {
	var p struct {
		Login    string
		Password string
		IP       string
	}

	err := internal.ReadBody(req, &p)
	if err != nil {
		internal.ForbiddenResponse(res, err, func() { logger.Errorf(err.Error()) })
	}

	logger.Infof("process request %v", p)

	found, checkResult, err := listsChecks(handler, p.IP)
	if err != nil {
		internal.OkResponse(res, false, fmt.Errorf("%v", err), func() { logger.Errorf(err.Error()) })
		return
	}

	if found {
		internal.OkWithDataResponse(res, fmt.Sprintf("ok=%v", checkResult))
		return
	}

	checkResult, err = mainChecks(handler, p.Login, p.Password, p.IP)
	if err != nil {
		internal.OkResponse(res, false, fmt.Errorf("%v", err), func() { logger.Errorf(err.Error()) })
		return
	}

	internal.OkWithDataResponse(res, fmt.Sprintf("ok=%v", checkResult))
}

//mainChecks in buckets
func mainChecks(handler *MainHandler, login, pass, addr string) (bool, error) {
	// get parameters from db
	params, err := handler.db.GetParametrs()
	if err != nil {
		return false, err
	}
	//check login
	loginRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("l_%s", login), params.N)
	if err != nil {
		return false, err
	}
	//check password
	passwordRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("p_%s", pass), params.M)
	if err != nil {
		return false, err
	}
	//check ip
	ipRes, err := checker.ProcessBucket(handler.store, fmt.Sprintf("i_%s", addr), params.K)
	if err != nil {
		return false, err
	}

	logger.Infof("login:%v=%v, pass:%v=%v, ip:%v=%v, result=>%v",
		login, loginRes, pass, passwordRes, addr, ipRes,
		loginRes && passwordRes && ipRes)

	return loginRes && passwordRes && ipRes, nil
}

//listsChecks in black/white lists
func listsChecks(handler *MainHandler, ip string) (bool, bool, error) {
	//black-white list ip check
	lists, err := handler.db.GetAddressList()
	if err != nil {
		logger.Errorf(err.Error())
		return false, false, err
	}

	for _, item := range *lists {
		result, err := checker.IsAddressInNewtork(item.Network, ip)
		if err != nil {
			return false, false, err
		}

		if result {
			logger.Infof("ip address %v contains in %v isWhite:%v", item.Network, ip, item.IsWhite)
			return true, item.IsWhite, nil
		}
	}

	return false, false, nil
}
