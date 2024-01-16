package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
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

// Note to self: Need to handle the case where the previous URL is null (when there has been no calls to map yet so no previous URL exits)
func GetPreviousPokeLocations(cfg *Config) error {
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
