package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/MPRaiden/pokedexcli/internal/pokeapi"
	"github.com/MPRaiden/pokedexcli/internal/pokecache"
	"time"
)
type cliCommand struct {
		name string
		description string
		callback func(*pokeapi.Config, *pokecache.Cache) error
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
			callback: func(cfg *pokeapi.Config, cache *pokecache.Cache) error {
                return pokeapi.GetPokeLocations(cfg, cache)
		},
	},
		"mapb": {
			name: "mapb",
			description: "Displays 20 previous pokemon locations",
			callback: func(cfg *pokeapi.Config, cache *pokecache.Cache) error {
		return pokeapi.GetPreviousPokeLocations(cfg, cache)
	    },
	},
	}

	func helpCommand(cfg *pokeapi.Config, cache *pokecache.Cache) error {
		fmt.Println("Welcome to the Pokedex!\n\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
		return nil
	}

	func exitCommand(cfg *pokeapi.Config, cache *pokecache.Cache) error {
		os.Exit(0)
		return nil
	}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)

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
			err := command.callback(config, cache)
			if err != nil {
				fmt.Printf("Failed to execute command %s: %s", command.name, err)
			}
		} else {
			fmt.Println("Your pokemon is >", text)
		}
	}
}
