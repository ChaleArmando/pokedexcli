package main

import (
	"fmt"
)

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
	}
	resp, err := c.client.GetLocationArea(c.previous)

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
