package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/MPRaiden/pokedexcli/pokeapi"
)
type cliCommand struct {
		name string
		description string
		callback func(*pokeapi.Config) error
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
		"map": {
			name: "map",
			description: "Displays 20 pokemon locations",
			callback: func(cfg *pokeapi.Config) error {
                return pokeapi.GetPokeLocations(cfg)
		},
	},
		"previousMap": {
			name: "previousMap",
			description: "Displays 20 previous pokemon locations",
			callback: func(cfg *pokeapi.Config) error {
		return pokeapi.GetPreviousPokeLocations(cfg)
	    },
	},
	}

	func helpCommand(cfg *pokeapi.Config) error {
		fmt.Println("Welcome to the Pokedex!\n\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
		return nil
	}

	func exitCommand(cfg *pokeapi.Config) error {
		os.Exit(0)
		return nil
	}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Initialize the config with initial PokeAPI URL
	config := &pokeapi.Config{
		Next: "https://pokeapi.co/api/v2/location/",
	}

	for {
		fmt.Print("pokedex> ")
		scanner.Scan()
		text := scanner.Text()

		command, exists := commands[text]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Printf("Failed to execute command %s: %s", command.name, err)
			}
		} else {
			fmt.Println("Your pokemon is >", text)
		}
	}
}
