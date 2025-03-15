package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/WST-T/Bobrdex/internal/pokeapi"
)

func startRepl() {
	cfg := &pokeapi.Config{CaughtPokemon: make(map[string]pokeapi.Pokemon)}
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' to see available commands")

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		text := reader.Text()

		if len(text) == 0 {
			fmt.Println("You entered nothing.")
			continue
		}

		cleaned := cleanInput(text)

		if len(cleaned) > 0 {
			commandName := cleaned[0]
			command, ok := commandRegistry[commandName]
			if !ok {
				fmt.Printf("Unknown command: %s\n", commandName)
				continue
			}

			args := []string{}
			if len(cleaned) > 1 {
				args = cleaned[1:]
			}

			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error executing command: %s\n", err)
			}
		} else {
			fmt.Println("You entered nothing.")
		}
		if reader.Err() != nil {
			fmt.Printf("Error reading input: %s\n", reader.Err())
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
