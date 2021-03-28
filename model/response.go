package model

type MatchResponse struct {
	TournamentID   int64   `json:"tid"`
	TournamentName string  `json:"name"`
	Matches        []Match `json:"matches"`
}
