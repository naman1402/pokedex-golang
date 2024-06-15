package main

import "fmt"

func Inspect(cfg *config, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]
	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("✅")
		fmt.Printf("\nName: %s \n", pokemon.Name)
		fmt.Printf("Height: %d \n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		fmt.Println("stats: ")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" 	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Printf("Types: ")
		for _, t := range pokemon.Types {
			fmt.Printf(" 	-%s\n", t.Type.Name)
		}

		return nil
	}

	fmt.Printf("you have not caught %s \n", name)
	return nil
}
