package main

import "fmt"

func Help(cfg *config, params []string) error {
	fmt.Println()
	fmt.Println("🎮 🕹️ 👾 Welcome to Pokedex")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
