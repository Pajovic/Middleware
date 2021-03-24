package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"middleware/handler"
	"middleware/model"
	"net/http"
)

type MatchService struct {
	httpHandler handler.HTTPHandler
}

func NewMatchService(httpHandler handler.HTTPHandler) MatchService {
	return MatchService{
		httpHandler: httpHandler,
	}
}

func (service *MatchService) GetMatches(url string) ([]model.Match, int, error) {
	resp, err := service.httpHandler.Get(url)
	if err != nil {
		log.Printf("[GetMatches] Failed retrieving data from endpoint: %s", err.Error())
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()
	var main model.MatchMain
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetMatches] ReadAll failed: %s", err.Error())
		return nil, http.StatusInternalServerError, err
	}
	if err = json.Unmarshal(body, &main); err != nil {
		log.Printf("[GetMatches] Error unmarshaling match: %s", err.Error())
		return nil, http.StatusOK, nil
	}
	if len(main.Doc) == 0 {
		errMsg := "[GetMatches] Missing doc"
		log.Print(errMsg)
		return nil, http.StatusInternalServerError, fmt.Errorf(errMsg)
	}

	var matches []model.Match
	for _, match := range main.Doc[0].Data.Matches {
		matches = append(matches, match)
	}

	return matches, http.StatusOK, nil
}
