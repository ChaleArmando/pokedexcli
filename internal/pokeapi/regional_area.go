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

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	var locationArea LocationArea
	fmt.Println(url)
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return locationArea, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea, fmt.Errorf("failed to crate request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea, fmt.Errorf("http client failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return locationArea, fmt.Errorf("failed to get response body: %w", err)
	}

	if res.StatusCode > 299 {
		return locationArea, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &locationArea)

	if err != nil {
		return locationArea, err
	}
	c.cache.Add(url, body)
	return locationArea, nil
}
