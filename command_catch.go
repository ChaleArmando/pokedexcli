package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(c *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("need to add pokemon to try catching")
	}
	pokemon := params[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	//Modify Function
	resp, err := c.client.GetPokemon(url)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	//Calculate if Pokemon was catched
	baseExp := resp.BaseExperience
	if rand.IntN(baseExp) > baseExp/2 {
		c.catchedPokemon[pokemon] = resp
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
