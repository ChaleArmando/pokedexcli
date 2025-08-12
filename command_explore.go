package main

import (
	"fmt"
)

func commandExplore(c *config, area string) error {
	if area == "" {
		return fmt.Errorf("need to add area to explore")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + area

	resp, err := c.client.GetLocationPokemon(url)

	if err != nil {
		return err
	}

	for _, result := range resp.PokemonEncounters {
		// Change, need to add new struct
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}
