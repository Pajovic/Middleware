package model

type MatchMain struct {
	QueryURL string     `json:"queryUrl"`
	Doc      []MatchDoc `json:"doc"`
}

type MatchDoc struct {
	Event string    `json:"event"`
	Data  MatchData `json:"data"`
}

type MatchData struct {
	Matches map[int]Match `json:"matches"`
}

type Match struct {
	ID       int             `json:"_id"`
	PlayTime PlayTime        `json:"time"`
	Teams    map[string]Team `json:"teams"`
	Comment  string          `json:"comment"`
}

type Team struct {
	ID         int    `json:"_id"`
	SID        int    `json:"_sid"`
	UID        int    `json:"uid"`
	Virtual    bool   `json:"virtual"`
	Name       string `json:"name"`
	MediumName string `json:"mediumname"`
	Abbr       string `json:"abbr"`
	Nickname   string `json:"nickname"`
	IsCountry  bool   `json:"iscountry"`
	HasLogo    bool   `json:"haslogo"`
}

type PlayTime struct {
	Time string `json:"time"`
	Date string `json:"date"`
	TZ   string `json:"tz"`
	UTS  int    `json:"uts"`
}
