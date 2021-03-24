package controller

import (
	"encoding/json"
	"fmt"
	"middleware/config"
	"middleware/model"
	"middleware/service"
	"net/http"
	"sort"

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
	}
	rsid, ok := vars["rcid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing real category id"))
		return
	}
	url := fmt.Sprintf("%s/%s/%s", config.Conf.ConfigTournamentsPath, sid, rsid)
	tournaments, status, err := controller.tournamentService.GetTournaments(url)
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
		return
	}
	for _, tournament := range tournaments {
		url := fmt.Sprintf("%s/%d/%d", config.Conf.FixturesTournamentPath, tournament.ID, 2021)
		matches, status, err := controller.matchService.GetMatches(url)
		if err != nil {
			w.WriteHeader(status)
			w.Write([]byte(err.Error()))
			return
		}
		if len(matches) > 0 {
			topMatches := SortMatches(matches, 5)
			response := model.MatchResponse{
				TournamentID: tournament.ID,
				Matches:      topMatches,
			}
			json.NewEncoder(w).Encode(response)
		}
	}
}

const layout = "01/02/06"

func SortMatches(matches []model.Match, num int) []model.Match {
	sortedTopMatches := make([]model.Match, num)
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].PlayTime.UTS > matches[j].PlayTime.UTS
	})

	for _, match := range matches {
		if match.ID == 0 {
			continue
		}
		sortedTopMatches = append(sortedTopMatches, match)
		if len(sortedTopMatches) == num {
			break
		}
		// t, err := time.Parse(layout, match.PlayTime.Date)
		// if err != nil {
		// 	log.Printf("err: %s", err.Error())
		// 	continue
		// }
		// if t.After(time.Now()) {
		// 	log.Printf("Match not yet held")
		// 	continue
		// }
		// sortedTopMatches = append(sortedTopMatches, match)
		// if len(sortedTopMatches) == num {
		// 	break
		// }
	}
	return sortedTopMatches
}
