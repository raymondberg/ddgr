package pkg

type SearchQuery struct {
	Filter SearchFilter `json:"filter"`
}

type SearchFilter struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Query string `json:"query"`
}
