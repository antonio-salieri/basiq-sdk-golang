package Utilities

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	host    string
	headers map[string]string
}

func NewAPI(host string) *API {
	return &API{
		host: host,
	}
}

func (api *API) Send(method string, path string, data []byte) ([]byte, int, error) {
	log.Println("Requesting: " + api.host + path)
	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, api.host+path, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(method, api.host+path, nil)
	}

	c := http.Client{}
	if err != nil {
		return nil, 0, err
	}

	for k, v := range api.headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func (api *API) SetHeader(header string, value string) *API {
	if api.headers == nil {
		api.headers = make(map[string]string)
	}
	api.headers[header] = value

	return api
}
