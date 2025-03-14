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
			// Just print the first word as the command
			fmt.Printf("Your command was: %s\n", cleaned[0])
		} else {
			fmt.Println("No valid command found.")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
