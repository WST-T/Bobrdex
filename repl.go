package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		text := reader.Text()

		if len(text) == 0 {
			fmt.Println("You entered nothing.")
			continue
		}

		// Clean the user input
		cleaned := cleanInput(text)

		// Check if there is at least one command
		if len(cleaned) > 0 {
			commandName := cleaned[0]
			command, ok := commandRegistry[commandName]
			if !ok {
				fmt.Printf("Unknown command: %s\n", commandName)
				continue
			}
			err := command.callback()
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
