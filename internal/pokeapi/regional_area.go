package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error) {
	var locationArea LocationArea

	res, err := http.Get(url)
	if err != nil {
		return locationArea, fmt.Errorf("failed to connect to pokemon api")
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return locationArea, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return locationArea, fmt.Errorf("failed to get response body: %w", err)
	}

	err = json.Unmarshal(body, &locationArea)

	if err != nil {
		return locationArea, err
	}
	return locationArea, nil
}
