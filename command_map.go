package main

import (
	"fmt"
)

func commandMap(c *config, params ...string) error {
	resp, err := c.client.GetLocationArea(c.next)

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
