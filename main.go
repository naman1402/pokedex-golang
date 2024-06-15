package main

import (
	"pokedex-go/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeApiClient: client,
	}

	startRepl(cfg)
}
