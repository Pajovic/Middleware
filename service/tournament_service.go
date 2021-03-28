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

type TournamentService struct {
	httpHandler handler.HTTPHandler
}

func NewTournamentService(httpHandler handler.HTTPHandler) TournamentService {
	return TournamentService{
		httpHandler: httpHandler,
	}
}

func (service *TournamentService) GetTournaments(serviceID, realCategoryID string) ([]model.Tournament, int, error) {
	url := fmt.Sprintf("%s/%s/%s", config.Conf.ConfigTournamentsPath, serviceID, realCategoryID)
	resp, err := service.httpHandler.Get(url)
	if err != nil {
		log.Printf("[GetTournaments] Failed retrieving data from endpoint: %s", err.Error())
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()
	var main model.TournamentMain
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetTournaments] ReadAll failed: %s", err.Error())
		return nil, http.StatusInternalServerError, err
	}
	if err = json.Unmarshal(body, &main); err != nil {
		log.Printf("[GetTournaments] Error unmarshaling tournament %s/%s:  %s", serviceID, realCategoryID, err.Error())
		return nil, http.StatusInternalServerError, err
	}
	if len(main.Doc) == 0 {
		errMsg := fmt.Sprintf("[GetTournaments] Missing doc for tournament %s/%s", serviceID, realCategoryID)
		log.Print(errMsg)
		return nil, http.StatusNotFound, fmt.Errorf(errMsg)
	}

	tournaments := make([]model.Tournament, len(main.Doc[0].Data.Tournaments))
	for i, tournament := range main.Doc[0].Data.Tournaments {
		tournaments[i] = tournament
	}
	return tournaments, http.StatusOK, nil
}
