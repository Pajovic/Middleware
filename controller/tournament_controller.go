package controller

import (
	"encoding/json"
	"middleware/model"
	"middleware/service"
	"middleware/util"
	"net/http"

	"github.com/gorilla/mux"
)

type TournamentController struct {
	tournamentService service.TournamentService
	matchService      service.MatchService
}

func NewTournamentController(tournamentService service.TournamentService, matchService service.MatchService) TournamentController {
	return TournamentController{
		tournamentService: tournamentService,
		matchService:      matchService,
	}
}

func (controller *TournamentController) GetTopMatches(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, ok := vars["sid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing service id"))
		return
	}
	rsid, ok := vars["rcid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing real category id"))
		return
	}
	tournaments, status, err := controller.tournamentService.GetTournaments(sid, rsid)
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}
	for _, tournament := range tournaments {
		//1 tournament could be held during 2 years e.g 20/21
		years := util.GetYears(tournament.Year)
		if len(years) == 0 {
			continue
		}
		var allMatches []model.Match
		for _, year := range years {
			matches, status, err := controller.matchService.GetMatches(tournament.ID, year)
			if err != nil {
				w.WriteHeader(status)
				w.Write([]byte(err.Error()))
				return
			}
			allMatches = append(allMatches, matches...)
		}
		if len(allMatches) > 0 {
			topMatches := util.SortMatches(allMatches, 5)
			response := model.MatchResponse{
				TournamentID:   tournament.ID,
				TournamentName: tournament.Name,
				Matches:        topMatches,
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}
