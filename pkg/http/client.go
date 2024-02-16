package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/itsbradn/govalo/pkg/conf"
)

func newClient() *http.Client {
	return &http.Client{Transport: &http.Transport{DialTLS: dialTLS}}
}

func newRequest(method, url string, data interface{}) (*http.Request, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header = conf.GetAuthHeaders().Clone()
	req.Header.Set("Referer", req.URL.Host)

	return req, nil
}

func SendRequest(method, url string, data any) (*http.Response, error) {
	req, err := newRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	res, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}
	return res,nil
}
