package main

import (
	"bufio"
	"fmt"
	"github.com/MPRaiden/pokedexcli/internal/pokeapi"
	"github.com/MPRaiden/pokedexcli/internal/pokecache"
	"github.com/MPRaiden/pokedexcli/models"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, *pokecache.Cache, []string) error
}

var commands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    helpCommand,
	},
	"exit": {
		name:        "exit",
		description: "Exits the pokedex",
		callback:    exitCommand,
	},
	"map": {
		name:        "map",
		description: "Displays 20 pokemon locations",
		callback:    pokeapi.GetPokeLocations,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays 20 previous pokemon locations",
		callback:    pokeapi.GetPreviousPokeLocations,
	},
	"explore": {
		name:        "explore",
		description: "Displays list of pokemon in a given location",
		callback:    pokeapi.GetPokeInLocation,
	},
	"catch": {
		name:        "catch",
		description: "Catches a pokemon",
		callback:    pokeapi.CatchPokemon,
	},
	"inspect": {
		name:        "inspect",
		description: "Displays information about a pokemon",
		callback:    inspectPokemon,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Displays all pokemon caught",
		callback:    displayPokedex,
	},
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 pokemon locations",
			callback:    pokeapi.GetPokeLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous pokemon locations",
			callback:    pokeapi.GetPreviousPokeLocations,
		},
		"explore": {
			name:        "explore",
			description: "Displays list of pokemon in a given location",
			callback:    pokeapi.GetPokeInLocation,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon",
			callback:    pokeapi.CatchPokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about a pokemon",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all pokemon caught",
			callback:    displayPokedex,
		},
	}
}

func helpCommand(cfg *pokeapi.Config, cache *pokecache.Cache, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func exitCommand(cfg *pokeapi.Config, cache *pokecache.Cache, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func inspectPokemon(cfg *pokeapi.Config, cache *pokecache.Cache, args []string) error {
	// Check if the user provided a pokemon name
	if len(args) < 1 {
		return fmt.Errorf("Please provide a pokemon name")
	}

	// Extract the pokemon name from the arguments and check if the pokemon is in the pokedex
	pokemonName := args[0]
	pokemon, ok := cfg.Trainer.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("You do not have a pokemon named %s", pokemonName)
	}

	// Print the pokemon information
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", int(pokemon.Height))
	fmt.Printf("Weight: %d\n", int(pokemon.Weight))
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		if stat.Name == "hp" || stat.Name == "attack" || stat.Name == "defense" || stat.Name == "special-attack" || stat.Name == "special-defense" || stat.Name == "speed" {
			fmt.Printf("\t-%s: %d\n", stat.Name, stat.BaseStat) // print only Base Stat, rename label
		}
	}
	fmt.Println("Types:")
	for _, type_ := range pokemon.Types {
		fmt.Printf("\tName: %s\n", type_.Name)
	}

	return nil
}

func displayPokedex(cfg *pokeapi.Config, cache *pokecache.Cache, args []string) error {
	// Displays all pokemon caught from pokedex (only names)
	if len(cfg.Trainer.Pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet")
		return nil
	}

	fmt.Println("Pokedex:")
	for _, pokemon := range cfg.Trainer.Pokedex {
		fmt.Printf("\t-%s\n", pokemon.Name)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Minute)

	// Initialize the config with initial PokeAPI URL
	config := &pokeapi.Config{
		Next:    "https://pokeapi.co/api/v2/location-area/",
		Trainer: &models.Trainer{Pokedex: make(map[string]models.Pokemon)},
	}

	fmt.Println("Welcome to the Pokedex!")
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
			fmt.Println("Unknown command.")
		}
	}
}
