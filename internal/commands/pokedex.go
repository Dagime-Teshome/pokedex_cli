package commands

import (
	"fmt"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func Pokedex(conf *shared.Config) error {
	if len(conf.PokeDex) <= 0 {
		fmt.Println("No pokemon in pokedex")
		return nil
	}
	fmt.Println("Your Pokedex: ")
	for _, value := range conf.PokeDex {
		fmt.Printf(" - %s \n", value.Name)
	}
	return nil
}
