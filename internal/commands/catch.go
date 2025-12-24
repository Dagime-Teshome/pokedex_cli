package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func Catch(conf *shared.Config) error {
	if len(conf.Data) <= 0 {
		return fmt.Errorf("No Pokemon found")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", conf.Data)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data:%s", err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error reading body:%s", err)
	}
	if resp.StatusCode >= 400 {
		fmt.Println("Pokemon not found (may not exist)")
		return nil
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return fmt.Errorf("Error marshaling :%s", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if simulateCatch(pokemon.BaseExperience) {
		value, ok := conf.PokeDex[pokemon.Name]
		if ok {
			fmt.Printf("%s already Caught\n", value.Name)
			return nil
		}
		conf.PokeDex[pokemon.Name] = shared.Pokemon(pokemon)
		fmt.Printf("%s was caught\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s has escaped\n", pokemon.Name)
	return nil
}

func simulateCatch(bXp int) bool {
	chance := 90 - bXp/3 // percentage

	if chance < 5 {
		chance = 5
	}

	roll := rand.Intn(100) // 0–99
	return roll < chance
}
