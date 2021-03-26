package handler

import (
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
	return http.Get(url)
}
