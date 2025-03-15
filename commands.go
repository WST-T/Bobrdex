package main

import (
	"fmt"
	"os"

	"github.com/WST-T/Bobrdex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config) error
}

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display names of 20 previous location areas",
			callback:    commandMapB,
		},
	}
}

func commandHelp(cfg *pokeapi.Config) error {
	fmt.Println("Usage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *pokeapi.Config) error {
	resp, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapB(cfg *pokeapi.Config) error {
	if cfg.Previous == "" {
		fmt.Println("No previous location areas.")
		return nil
	}

	resp, err := pokeapi.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = resp.Next
	cfg.Previous = resp.Previous

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
