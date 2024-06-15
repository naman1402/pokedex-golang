package main

import "fmt"

func Pokedex(cfg *config, args []string) error {
	fmt.Println("🎴 Your pokedex: ")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" -%s\n", pokemon)
	}

	return nil
}
