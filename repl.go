package main

import (
	"strings"
)

func cleanInput(text string) []string {
	inputs := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return inputs
}
