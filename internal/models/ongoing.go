package models

type Ongoing struct {
	OriginalTitle string `json:"original_title"`
	RussianTitle  string `json:"russian_title"`
	Link          string `json:"link"`
}

type OngoingResponse struct {
	Day       string    `json:"day"`
	DaysAhead int       `json:"days_ahead"`
	Titles    []Ongoing `json:"titles"`
}
