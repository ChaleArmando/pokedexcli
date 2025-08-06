package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ChaleArmando/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	inputs := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return inputs
}

func replLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	configPokeAPI := &config{
		client:   pokeapi.NewClient(time.Second*5, time.Minute*5),
		next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
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
