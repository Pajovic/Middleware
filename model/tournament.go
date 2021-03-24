package model

type TournamentMain struct {
	QueryURL string          `json:"queryUrl"`
	Doc      []TournamentDoc `json:"doc"`
}

type TournamentDoc struct {
	Event string         `json:"event"`
	Data  TournamentData `json:"data"`
}

type TournamentData struct {
	Tournaments []Tournament `json:"tournaments"`
}

type Tournament struct {
	ID   int    `json:"_id"`
	Name string `json:"name"`
	Year string `json:"year"`
}
