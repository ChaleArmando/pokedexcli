package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	var pokemon Pokemon

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return pokemon, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, fmt.Errorf("http client failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return pokemon, fmt.Errorf("failed to get response body: %w", err)
	}

	if res.StatusCode > 299 {
		return pokemon, fmt.Errorf("response failed with status code: %d\nbody: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return pokemon, err
	}
	c.cache.Add(url, body)
	return pokemon, nil
}
