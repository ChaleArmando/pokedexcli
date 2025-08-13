package main

import (
	"fmt"
)

func commandExplore(c *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("need to add area to explore")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + params[0]

	resp, err := c.client.GetLocationPokemon(url)

	if err != nil {
		return err
	}

	for _, result := range resp.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}
