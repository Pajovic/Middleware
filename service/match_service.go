package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"middleware/config"
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

func (service *MatchService) GetMatches(tournamentID, year int64) ([]model.Match, int, error) {
	url := fmt.Sprintf("%s/%d/%d", config.Conf.FixturesTournamentPath, tournamentID, year)
	resp, err := service.httpHandler.Get(url)
	if err != nil {
		log.Printf("[GetMatches] Failed retrieving data from endpoint: %s", err.Error())
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()
	var main model.MatchMain
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetMatches] ReadAll failed for tournament %d: %s", tournamentID, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	if err = json.Unmarshal(body, &main); err != nil {
		//Common case is that unmarshall failed because there are no matches related to this tournament
		// and matches struct is empty.
		log.Printf("[GetMatches] Error unmarshaling tournament %d match: %s", tournamentID, err.Error())
		return nil, http.StatusOK, nil
	}
	if len(main.Doc) == 0 {
		errMsg := fmt.Sprintf("[GetMatches] Missing doc for tournament: %d", tournamentID)
		log.Print(errMsg)
		return nil, http.StatusInternalServerError, fmt.Errorf(errMsg)
	}

	var matches []model.Match
	for _, match := range main.Doc[0].Data.Matches {
		matches = append(matches, match)
	}

	return matches, http.StatusOK, nil
}
