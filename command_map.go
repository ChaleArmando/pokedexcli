package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandMap(c *config) error {
	resp, err := pokeapi.GetLocationArea(c.next)

	if err != nil {
		return err
	}

	c.next = resp.Next
	c.previous = resp.Previous

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
