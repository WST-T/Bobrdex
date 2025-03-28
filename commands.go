package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"

	"github.com/WST-T/Bobrdex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, args ...string) error
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
		"explore": {
			name:        "explore <location area>",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "View information about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func commandHelp(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Usage:")
	for _, cmd := range commandRegistry {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *pokeapi.Config, args ...string) error {
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

func commandMapB(cfg *pokeapi.Config, args ...string) error {
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

func commandExplore(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area name.")
		return nil
	}

	locationAreaName := args[0]
	fmt.Printf("Exploring location area: %s\n", locationAreaName)

	locationArea, err := pokeapi.GetLocationAreaDetails(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon encounters:")
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this area.")
	} else {
		for _, encounter := range locationArea.PokemonEncounters {
			fmt.Printf("  %s\n", encounter.Pokemon.Name)
		}
	}

	return nil
}

func commandCatch(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Pokemon name.")
		return nil
	}

	pokemonName := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := pokeapi.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	catchChance := 0.0
	if pokemon.BaseExperience <= 0 {
		catchChance = 0.7
	} else {
		catchChance = 1.0 - float64(pokemon.BaseExperience)/1000.0
		if catchChance < 0.3 {
			catchChance = 0.3
		}
	}

	randValue := rand.Float64()
	caught := randValue <= catchChance
	if caught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s broke free!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *pokeapi.Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokemon name")
	}

	pokemonName := strings.ToLower(args[0])

	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("You haven't caught %s yet.\n", pokemonName)
		return nil
	}

	fmt.Printf("Pokemon: %s\n", pokemon.Name)
	fmt.Printf("Base experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *pokeapi.Config, args ...string) error {
	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}

	fmt.Println("Caught Pokemon:")

	var names []string
	for name := range cfg.CaughtPokemon {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		pokemon := cfg.CaughtPokemon[name]
		fmt.Printf("  %s\n", pokemon.Name)

		for i, t := range pokemon.Types {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(t.Type.Name)
		}
		fmt.Println()
	}

	fmt.Printf("Total caught: %d\n", len(cfg.CaughtPokemon))

	return nil
}
