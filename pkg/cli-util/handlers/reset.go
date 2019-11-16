package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Reset(server, address, login string) (*RespData, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"login":"%s", "ip":"%s"}`, login, address)).
		Post(fmt.Sprintf("%s/admin/reset", server))
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
