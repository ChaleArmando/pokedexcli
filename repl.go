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
	configPokeAPI := &config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInputSlice := cleanInput(input)
		if len(cleanInputSlice) == 0 {
			continue
		}

		if val, ok := getCommands()[cleanInputSlice[0]]; ok {

			val.callback(configPokeAPI)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
