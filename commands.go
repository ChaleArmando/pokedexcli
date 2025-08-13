package main

import (
	"github.com/ChaleArmando/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	client         pokeapi.Client
	next           string
	previous       string
	catchedPokemon map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	commands_map := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémons in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try catching Pokemons by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon previously catched",
			callback:    commandInspect,
		},
	}
	return commands_map
}
