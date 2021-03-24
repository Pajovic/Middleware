package handler

import (
	"log"
	"net/http"
)

type HTTPHandler interface {
	Get(string) (*http.Response, error)
}

type httpHandler struct {
}

func NewHttpHandler() HTTPHandler {
	return &httpHandler{}
}

func (handler *httpHandler) Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[httpHandler Get] Failed retrieving data from endpoint: %s", err.Error())
		return nil, err
	}
	return resp, nil
}
