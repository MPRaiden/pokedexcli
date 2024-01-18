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

func GetPokeLocations(cfg *Config, cache *pokecache.Cache) error {
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

func GetPreviousPokeLocations(cfg *Config, cache *pokecache.Cache) error {
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
