package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex-go/pokeapi"
	"strings"
)

type config struct {
	pokeApiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
	caughtPokemon map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		parameters := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, parameters)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(txt string) []string {
	output := strings.ToLower(txt)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Help,
		},
		"map": {
			name:        "map",
			description: "Get next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of locations",
			callback:    commandMap2,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    Explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a Pokemon",
			callback:    Catch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a Pokemon",
			callback:    Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists caught Pokemon",
			callback:    Pokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    Exit,
		},
	}
}
