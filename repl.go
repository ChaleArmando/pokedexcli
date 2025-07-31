package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	inputs := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return inputs
}

func replLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInputSlice := cleanInput(input)
		if len(cleanInputSlice) == 0 {
			continue
		}

		if val, ok := getCommands()[cleanInputSlice[0]]; ok {
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
