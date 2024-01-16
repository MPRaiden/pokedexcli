package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPokeLocations() {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("HTTP Error: %s", body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
