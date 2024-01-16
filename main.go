package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/MPRaiden/pokedex/pokedexcli/pokeapi"
)
type cliCommand struct {
		name string
		description string
		callback func() error
	}

type config struct {
	Next string
	Previous string
}

	var commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: helpCommand,
		},
		"exit": {
			name: "exit",
			description: "Exits the pokedex",
			callback: exitCommand,
		},
		"map": {}.
		"mapb": {}.
	}

	func helpCommand() error {
		fmt.Println("Welcome to the Pokedex!\n\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
		return nil
	}

	func exitCommand() error {
		os.Exit(0)
		return nil
	}

func main() {
	scanner := bufio.NewScanner(os.Stdin):

	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		text := scanner.Text()

		command, exists := commands[text]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Printf("Failed to execute command %s: %s", command.name, err)
			}
		} else {
			fmt.Println("Your pokemon is >", text)
			pokeapi.GetPokeLocations()
		}
	}
}
