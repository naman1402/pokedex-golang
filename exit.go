package main

import (
	"fmt"
	"os"
)

func Exit(cfg *config, params []string) error {
	fmt.Println("🦇 Exiting pokedex")
	os.Exit(0)
	return nil
}
