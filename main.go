package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/MPRaiden/pokedexcli/internal/pokeapi"
	"github.com/MPRaiden/pokedexcli/internal/pokecache"
	"time"
	"strings"
	"github.com/MPRaiden/pokedexcli/models"
)

type cliCommand struct {
		name string
		description string
		callback func(*pokeapi.Config, *pokecache.Cache, []string) error
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
			callback: pokeapi.GetPokeLocations,
		},
		"mapb": {
			name: "mapb",
			description: "Displays 20 previous pokemon locations",
			callback: pokeapi.GetPreviousPokeLocations,
		},	
		"explore": {
			name: "explore",
			description: "Displays list of pokemon in a given location",
			callback: pokeapi.GetPokeInLocation,
		},		
		"catch": {
			name: "catch",
			description: "Catches a pokemon",
			callback: pokeapi.CatchPokemon,
		},
}

func helpCommand(cfg *pokeapi.Config, cache *pokecache.Cache, args[]string) error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\nmap: Displays 20 pokemon locations\nmapb: Displays 20 previous pokemon locations\nexplore: Displays list of pokemon in a given location\ncatch: Attempts to catch a pokemon and if successful saves it to players pokedex.")
		return nil
	}

func exitCommand(cfg *pokeapi.Config, cache *pokecache.Cache, args[]string) error {
		os.Exit(0)
		return nil
	}	

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)

	// Initialize the config with initial PokeAPI URL
	config := &pokeapi.Config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Trainer: &models.Trainer{Pokedex: make(map[string]models.Pokemon)},
	}

	for {
		fmt.Print("pokedex> ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "reading standard input: %v\n", err)
			}
			break
		}
		input := scanner.Text()
		// Split the user input into command and arguments
		parts := strings.Split(input, " ")

		commandName := parts[0]
		args := parts[1:]

		command, exists := commands[commandName]
		if exists {
			err := command.callback(config, cache, args)
			if err != nil {
				fmt.Printf("Failed to execute command %s: %s", command.name, err)
			}
		} else {
			fmt.Println("Your input is >", input)
		}
	}
}
