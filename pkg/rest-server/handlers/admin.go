package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func (handler *MainHandler) ResetBucket(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error read request body", http.StatusForbidden)
		return
	}
	var r struct {
		Login string
		IP    string
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error parsing body", http.StatusForbidden)
		return
	}
	logger.Infof("process ResetBucket %v", r)
	err = handler.store.RemoveBuckets(fmt.Sprintf("l_%s", r.Login), fmt.Sprintf("i_%s", r.IP))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "bucket remove failed", http.StatusForbidden)
		return
	}
	//nolint
	res.Write([]byte("ok"))
}

func (handler *MainHandler) AddToWhiteList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *MainHandler) RemoveFromWhiteList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *MainHandler) AddToBlackList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *MainHandler) RemoveFromBlackList(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("ok"))
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *MainHandler) UpdateParameters(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error read request body", http.StatusForbidden)
		return
	}
	var p struct {
		K, M, N int
	}
	err = json.Unmarshal(body, &p)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "error parsing body", http.StatusForbidden)
		return
	}
	logger.Infof("process update parameters %v", p)

	err = handler.db.UpdateParametrs(p.K, p.M, p.N)
	if err != nil {
		logger.Errorf(err.Error())
		http.Error(res, "update parameters failed", http.StatusForbidden)
		return
	}
	//nolint
	res.Write([]byte("ok"))
}
