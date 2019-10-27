package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//nolint
type respData struct {
	Success     bool        `json:"success"`
	Data        interface{} `json:"data,omitempty"`
	Description string      `json:"description,omitempty"`
}

// ReadBody and serializa to interface
func ReadBody(req *http.Request, p interface{}) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("read body error %v", err)
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		return fmt.Errorf("parsing body error %v", err)
	}

	return nil
}

func OkResponse(res http.ResponseWriter, success bool, err error, callback func()) {
	r := respData{Success: success}
	if err != nil {
		r.Description = err.Error()

		callback()
	}

	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(res, fmt.Sprintf("response marshal error %v", err), http.StatusInternalServerError)
		return
	}

	//nolint
	res.Write(resp)
}

func OkWithDataResponse(res http.ResponseWriter, data interface{}) {
	r := respData{Success: true, Data: data}

	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(res, fmt.Sprintf("response marshal error %v", err), http.StatusInternalServerError)
		return
	}

	//nolint
	res.Write(resp)
}

func ForbiddenResponse(res http.ResponseWriter, err error, callback func()) {
	r := respData{Success: false, Description: err.Error()}

	resp, err := json.Marshal(r)
	if err != nil {
		http.Error(res, fmt.Sprintf("response marshal error %v", err), http.StatusInternalServerError)
		return
	}

	http.Error(res, string(resp), http.StatusForbidden)
	callback()
}
