package main

import (
	"fmt"
	"math/rand"
)

func Catch(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeApiClient.GetPokemonByName(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("\n⚔️  Throwing a pokeball at %s ... \n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped! 📢\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught🏆\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil

}
