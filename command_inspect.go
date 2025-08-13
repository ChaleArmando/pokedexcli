package main

import (
	"fmt"
)

func commandInspect(c *config, params ...string) error {
	if len(params) != 1 {
		return fmt.Errorf("need to add pokemon to inspect")
	}

	if pokemon, ok := c.catchedPokemon[params[0]]; ok {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("Stats:")
		for _, entry := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", entry.Stat.Name, entry.BaseStat)
		}

		fmt.Println("Types:")
		for _, entry := range pokemon.Types {
			fmt.Printf(" - %s\n", entry.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
