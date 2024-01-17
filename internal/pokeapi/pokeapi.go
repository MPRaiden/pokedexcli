package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"errors"
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

func GetPokeLocations(cfg *Config) error {
	res, err := http.Get(cfg.Next)
	if err != nil {
		return fmt.Errorf("Failed to get locations: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %w", err)
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("HTTP Error: %s", body)
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

func GetPreviousPokeLocations(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("There are no previous locations")
		return errors.New("There are no previous locations to display")
	}
	res, err := http.Get(cfg.Previous)
	if err != nil {
		return fmt.Errorf("Failed to get locations: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Failed to read response body: %w", err)
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("HTTP Error: %s", body)
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
