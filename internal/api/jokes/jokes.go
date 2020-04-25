package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"workshop/internal/api"
)

const jokePath = "/api?format=json"

type JokeClient struct {
	url string
}

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + jokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)

	return &data, err
}
