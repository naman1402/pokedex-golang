package main

import (
	"errors"
	"fmt"
)

type Config struct {
	Previous string
	Next     string
}

func commandMap(cfg *config, params []string) error {
	locations, err := cfg.pokeApiClient.ListLocations(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous
	fmt.Println()
	for _, location := range locations.Results {
		fmt.Println(location.Name)
		fmt.Println()
	}
	return nil
}

func commandMap2(cfg *config, params []string) error {
	if cfg.prevLocation == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeApiClient.ListLocations(cfg.prevLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.prevLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
