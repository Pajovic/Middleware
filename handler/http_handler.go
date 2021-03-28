package handler

import (
	"net/http"
)

type HTTPHandler interface {
	Get(string) (*http.Response, error)
}

type httpHandler struct {
	httpClient *http.Client
}

func NewHttpHandler() HTTPHandler {
	return &httpHandler{
		httpClient: &http.Client{},
	}
}

func (handler *httpHandler) Get(url string) (*http.Response, error) {
	return handler.httpClient.Get(url)
}
