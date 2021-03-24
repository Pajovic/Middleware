package model

type MatchResponse struct {
	TournamentID int     `json:"tid"`
	Matches      []Match `json:"matches"`
}
