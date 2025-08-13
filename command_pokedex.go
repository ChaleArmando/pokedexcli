package main

import (
	"fmt"
)

func commandPokedex(c *config, params ...string) error {
	fmt.Println("Your Pokedex: ")

	for key := range c.catchedPokemon {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
