package pkg

type SearchResult struct {
	Data []SearchResultRow `json:"data"`
}

type SearchResultRow struct {
	Attributes SearchAttributes `json:"attributes"`
}
type SearchAttributes struct {
	Timestamp string   `json:"timestamp"`
	Service   string   `json:"service"`
	Message   string   `json:"message"`
	Tags      []string `json:"tags"`
}
