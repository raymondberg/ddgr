package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Search(query string, config Config) SearchResult {
	search := SearchQuery{
		Filter: SearchFilter{
			From:  "2023-01-10T00:00:00+00:00",
			To:    "2023-01-11T23:25:00+00:00",
			Query: query,
		},
	}

	return SearchQuerySearch(search, config)
}

func SearchQuerySearch(search SearchQuery, config Config) SearchResult {
	data, err := json.Marshal(search)
	if err != nil {
		println("Invalid json structure")
		panic(err)
	}

	request, _ := http.NewRequest(
		"POST",
		"https://api.datadoghq.com/api/v2/logs/events/search",
		bytes.NewReader(data),
	)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("DD-API-KEY", config.apikey)
	request.Header.Set("DD-APPLICATION-KEY", config.appkey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var result SearchResult

	if err := json.Unmarshal(body, &result); err != nil {
		println(err)
		panic(err)
	}
	return result
}
