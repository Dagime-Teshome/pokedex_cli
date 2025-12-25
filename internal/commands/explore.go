package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func Explore(conf *shared.Config) error {
	if len(conf.Data) <= 0 {
		return fmt.Errorf("Empty location not allowed")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", conf.Data)

	value, ok := conf.Cache.Get(url)

	if ok {
		location_area := shared.LocationArea{}
		marshErr := json.Unmarshal(value, &location_area)
		if marshErr != nil {
			return fmt.Errorf("error marshalling: Error:%s", marshErr)
		}
		printPokemon(conf.Data, location_area)
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch area data %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	conf.Cache.Add(url, body)
	location_area := shared.LocationArea{}
	marshErr := json.Unmarshal(body, &location_area)
	if marshErr != nil {
		return fmt.Errorf("error marshalling: Error:%s", marshErr)
	}
	printPokemon(conf.Data, location_area)
	return nil
}

func printPokemon(area string, locationArea shared.LocationArea) {
	fmt.Println("Exploring......")
	fmt.Printf("Found this pokemon in %s : \n", area)
	for _, value := range locationArea.Pokemon_Encounters {
		fmt.Printf(" - %s\n", value.Pokemon.Name)
	}
}
