package main

import (
	"fmt"
)

func commandMapb(c *config, params ...string) error {
	if c.previous == "" {
		return fmt.Errorf("you're on the first page")
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
