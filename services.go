package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Service struct {
	endpoint string
}

func CreateService(endpoint string) Service {
	return Service{
		endpoint: endpoint,
	}
}

func (service *Service) Get(path string, cookie string) (*http.Response, error) {
	return service.Request("GET", path, cookie, nil)
}

func (service *Service) Post(path string, cookie string, body map[string]string) (*http.Response, error) {
	return service.Request("POST", path, cookie, body)
}

func (service *Service) Put(path string, cookie string, body map[string]string) (*http.Response, error) {
	return service.Request("PUT", path, cookie, body)
}

func (service *Service) Delete(path string, cookie string, body map[string]string) (*http.Response, error) {
	return service.Request("DELETE", path, cookie, body)
}

func (service *Service) Request(method string, path string, cookie string, body map[string]string) (*http.Response, error) {
	var data io.Reader = nil

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, service.endpoint+path, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if cookie != "" {

		req.Header.Add("Cookie", "jwt="+cookie)
	}
	client := &http.Client{}
	return client.Do(req)
}
