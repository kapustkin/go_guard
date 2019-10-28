package handlers

import (
	"fmt"
	"net/http"

	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal"
	"github.com/kapustkin/go_guard/pkg/rest-server/handlers/internal/admin"
	logger "github.com/sirupsen/logrus"
)

func (handler *MainHandler) ResetBucket(res http.ResponseWriter, req *http.Request) {
	var p struct {
		Login string
		IP    string
	}

	err := internal.ReadBody(req, &p)
	if err != nil {
		internal.ForbiddenResponse(res, err, func() { logger.Errorf(err.Error()) })
		return
	}

	logger.Infof("process ResetBucket %v", p)
	err = handler.store.RemoveBuckets(fmt.Sprintf("l_%s", p.Login), fmt.Sprintf("i_%s", p.IP))

	if err != nil {
		internal.OkResponse(res, false, err, func() { logger.Errorf(err.Error()) })
		return
	}

	internal.OkResponse(res, true, nil, nil)
}

func (handler *MainHandler) GetAllLists(res http.ResponseWriter, req *http.Request) {
	list, err := handler.db.GetAddressList()
	if err != nil {
		internal.OkResponse(res, false, err, func() { logger.Errorf(err.Error()) })
		return
	}

	type rec struct {
		IsWhite bool
		Network string
	}

	result := make([]rec, len(*list))
	for i, item := range *list {
		result[i] = rec{IsWhite: item.IsWhite, Network: item.Network}
	}

	internal.OkWithDataResponse(res, &result)
}

func (handler *MainHandler) AddToList(res http.ResponseWriter, req *http.Request) {
	var p struct {
		Network string
		IsWhite bool
	}

	err := internal.ReadBody(req, &p)
	if err != nil {
		internal.ForbiddenResponse(res, err, func() { logger.Errorf(err.Error()) })
		return
	}

	err = admin.IsAddressValid(p.Network)
	if err != nil {
		internal.OkResponse(res, false, err, func() { logger.Errorf(err.Error()) })
		return
	}

	logger.Infof("add to list %v", p)

	err = handler.db.AddAddress(p.Network, p.IsWhite)
	if err != nil {
		internal.OkResponse(res, false, fmt.Errorf("%v", err), func() { logger.Errorf(err.Error()) })
		return
	}

	internal.OkResponse(res, true, nil, nil)
}

func (handler *MainHandler) RemoveFromList(res http.ResponseWriter, req *http.Request) {
	var p struct {
		Network string
		IsWhite bool
	}

	err := internal.ReadBody(req, &p)
	if err != nil {
		internal.ForbiddenResponse(res, err, func() { logger.Errorf(err.Error()) })
		return
	}

	logger.Infof("remove from list %v", p)
	// проверка валидности адреса

	err = handler.db.RemoveAddress(p.Network, p.IsWhite)
	if err != nil {
		internal.OkResponse(res, false, fmt.Errorf("%v", err), func() { logger.Errorf(err.Error()) })
		return
	}

	internal.OkResponse(res, true, nil, nil)
}

func (handler *MainHandler) UpdateParameters(res http.ResponseWriter, req *http.Request) {
	var p struct {
		K, M, N int
	}

	err := internal.ReadBody(req, &p)
	if err != nil {
		internal.ForbiddenResponse(res, err, func() { logger.Errorf(err.Error()) })
		return
	}

	logger.Infof("process update parameters %v", p)

	err = handler.db.UpdateParametrs(p.K, p.M, p.N)
	if err != nil {
		internal.OkResponse(res, false, fmt.Errorf("%v", err), func() { logger.Errorf(err.Error()) })
		return
	}

	internal.OkResponse(res, true, nil, nil)
}
