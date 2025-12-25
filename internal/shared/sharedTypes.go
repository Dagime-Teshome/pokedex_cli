package shared

import (
	"github.com/Dagime-Teshome/pokedex_cli/internal/pokecache"
)

type Config struct {
	Previous string
	Next     string
	Cache    pokecache.Cache
	Data     string
	PokeDex  map[string]Pokemon
}

func (c *Config) SetPrev(s *string) {
	if s == nil {
		c.Previous = "null"
		return
	}
	c.Previous = *s
}
func (c *Config) SetNext(s *string) {
	if s == nil {
		c.Next = "null"
		return
	}
	c.Next = *s
}

type Locations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []result
}

type result struct {
	Name string
	Url  string
}

type Pokemon struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Url            string        `json:"url"`
	BaseExperience int           `json:"base_experience"`
	Stats          []Stat        `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type Stat struct {
	Stat      Name `json:"stat"`
	Effort    int  `json:"effort"`
	Base_Stat int  `json:"base_stat"`
}

type Name struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	Slot int  `json:"slot"`
	Type Name `json:"type"`
}

type LocationArea struct {
	Id                    int                   `json:"id"`
	Name                  string                `json:"name"`
	Game_index            int                   `json:"game_index"`
	Encounter_MethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Pokemon_Encounters    []PokemonEncounter    `json:"pokemon_encounters"`
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

// type Pokemon struct {
// 	Id             int    `json:"id"`
// 	Name           string `json:"name"`
// 	Url            string `json:"url"`
// 	BaseExperience int    `json:"base_experience"`
// }
