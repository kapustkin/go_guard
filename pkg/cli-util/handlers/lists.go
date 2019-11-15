package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetAllList(server string) (*RespData, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/admin/lists", server))
	if err != nil {
		return nil, err
	}

	var res RespData
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func AddToList(server, network string, isWhite bool) (*RespData, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"Network":"%s", "IsWhite":%v}`, network, isWhite)).
		Post(fmt.Sprintf("%s/admin/lists/add", server))
	if err != nil {
		return nil, err
	}

	var res RespData
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func RemoveFromList(server, network string, isWhite bool) (*RespData, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"Network":"%s", "IsWhite":%v}`, network, isWhite)).
		Post(fmt.Sprintf("%s/admin/lists/remove", server))
	if err != nil {
		return nil, err
	}

	var res RespData
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetParams(server string) (*RespParams, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s/admin/params", server))
	if err != nil {
		return nil, err
	}

	var res RespParams
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func SetParams(server string, k, m, n int32) (*RespParams, error) {
	return nil, fmt.Errorf("sdasd")
}
