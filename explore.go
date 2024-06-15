package main

import "fmt"

func Explore(cfg *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide an area name to explore")
	}

	areaName := params[0]
	pokemonNames, err := cfg.pokeApiClient.ExploreArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("\n🔎 Exploring %s .... \n🦖 Found Pokemon\n", areaName)
	for _, name := range pokemonNames {
		fmt.Printf(" 	-%s\n", name)
	}
	return nil
}
