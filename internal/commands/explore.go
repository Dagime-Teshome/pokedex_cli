package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

// id:1
// name:"canalave-city-area"
// game_index:1
type LocationArea struct {
	Id                    int                   `json:"id"`
	Name                  string                `json:"name"`
	Game_index            int                   `json:"game_index"`
	Encounter_MethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Pokemon_Encounters    []PokemonEncounter    `json:"pokemon_encounters"`
}
type Name struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type EncounterMethod struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Names Name   `json:"names"`
}
type EncounterVersionDetails struct {
	rate int
}

type EncounterMethodRate struct {
	Encounter_Method EncounterMethod `json:"encounter_method"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Url            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}

func Explore(conf *shared.Config) error {
	if len(conf.Data) <= 0 {
		return fmt.Errorf("Empty location not allowed")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", conf.Data)

	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("Couldn't fetch area data %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	location_area := LocationArea{}
	marshErr := json.Unmarshal(body, &location_area)
	if marshErr != nil {
		return fmt.Errorf("error marshalling: Error:%s", marshErr)
	}
	fmt.Println("Exploring......")
	fmt.Printf("Found this pokemon in %s : \n", conf.Data)
	for _, value := range location_area.Pokemon_Encounters {
		fmt.Printf(" - %s\n", value.Pokemon.Name)
	}
	return nil
}
