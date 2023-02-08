package pkg

type SearchResult struct {
	Data []SearchResultRow `json:"data"`
}

type SearchResultRow struct {
	Id         string           `json:"id"`
	Attributes SearchAttributes `json:"attributes"`
}

type SearchAttributes struct {
	Timestamp string   `json:"timestamp"`
	Service   string   `json:"service"`
	Message   string   `json:"message"`
	Tags      []string `json:"tags"`
}
