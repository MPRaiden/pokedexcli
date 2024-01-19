package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"errors"
	"github.com/MPRaiden/pokedexcli/internal/pokecache"
)

type Config struct {
	Next string
	Previous string
}

type LocationsResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

func GetPokeLocations(cfg *Config, cache *pokecache.Cache, args[]string) error {
	var err error
	body, ok := cache.Get(cfg.Next)
	if !ok {
		var res *http.Response
		res, err = http.Get(cfg.Next)
		if err != nil {
			return fmt.Errorf("Failed to get locations: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Failed to read response body: %w", err)
		}

		if res.StatusCode > 299 {
			return fmt.Errorf("HTTP Error: %s", body)
		}
	
		cache.Add(cfg.Next, body)
	}

	var locationsResponse LocationsResponse
	err = json.Unmarshal(body, &locationsResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	for _, result := range locationsResponse.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationsResponse.Next
	cfg.Previous = locationsResponse.Previous

	return nil
}

func GetPreviousPokeLocations(cfg *Config, cache *pokecache.Cache, args[]string) error {
	if cfg.Previous == "" {
		return errors.New("There are no previous locations to display")
	}
	var err error
	var res *http.Response
	var body []byte
	body, ok := cache.Get(cfg.Previous)
	if !ok {
		res, err = http.Get(cfg.Previous)
		if err != nil {
			return fmt.Errorf("Failed to get locations: %w", err)
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Failed to read response body: %w", err)
		}

		if res.StatusCode > 299 {
			return fmt.Errorf("HTTP Error: %s", body)
		}
	}

	var locationsResponse LocationsResponse
	err = json.Unmarshal(body, &locationsResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	for _, result := range locationsResponse.Results {
		fmt.Println(result.Name)
	}
	cfg.Next = locationsResponse.Next
	cfg.Previous = locationsResponse.Previous

	return nil
}

func GetPokeInLocation(cfg *Config, cache *pokecache.Cache, args[]string) error {
	if len(args) < 1 {
		return errors.New("Please provide a location")
	}
	location := args[0]

	// Check if the location is in the cache
	data, isCached := cache.Get(location)
	if isCached {
		fmt.Println("Using cached data")

		var pokemons []string
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return err
		}
		fmt.Println(pokemons)
	} else {
		resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299  || resp.StatusCode < 200 {
			return fmt.Errorf("HTTP Error: %s", resp.Status)
		}

		// Since JSON is basically a map where keys are strings but the values can be anything
		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}
		// Extract Pokemon names and print them
		var pokemonNames []string
		for _, pokemon := range result["pokemon_encounters"].([]interface{}) {
			pokemonNames = append(pokemonNames, pokemon.(map[string]interface{})["pokemon"].(map[string]interface{})["name"].(string))
		}
    
		dataToCache, err := json.Marshal(pokemonNames)
		if err != nil {
			return err
		}
		cache.Add(location, dataToCache) // Remember to cache the result after the API call
	
		fmt.Printf("Found Pokemon in %s: \n", location)
		for _, name := range pokemonNames {
			fmt.Printf(" - %s\n", name)
		}
	} 
	return nil
}
