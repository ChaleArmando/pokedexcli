package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	var locationArea LocationArea
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return locationArea, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea, fmt.Errorf("failed to create request: %w", err)
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
		return locationArea, fmt.Errorf("response failed with status code: %d\nbody: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &locationArea)

	if err != nil {
		return locationArea, err
	}
	c.cache.Add(url, body)
	return locationArea, nil
}

func (c *Client) GetLocationPokemon(url string) (LocationPokemon, error) {
	var locationPokemon LocationPokemon
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationPokemon)
		if err != nil {
			return locationPokemon, err
		}
		return locationPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationPokemon, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationPokemon, fmt.Errorf("http client failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return locationPokemon, fmt.Errorf("failed to get response body: %w", err)
	}

	if res.StatusCode > 299 {
		return locationPokemon, fmt.Errorf("response failed with status code: %d\nbody: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &locationPokemon)

	if err != nil {
		return locationPokemon, err
	}
	c.cache.Add(url, body)
	return locationPokemon, nil
}
