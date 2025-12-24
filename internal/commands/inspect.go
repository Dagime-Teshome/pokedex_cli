package commands

import (
	"fmt"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func Inspect(conf *shared.Config) error {
	if len(conf.Data) <= 0 {
		return fmt.Errorf("No inspect parameter found")
	}
	value, ok := conf.PokeDex[conf.Data]
	if !ok {
		fmt.Printf(" you have not caught that pokemon ,%s\n", conf.Data)
		return nil
	}
	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Url: %s\n", value.Url)
	fmt.Printf("Base XP: %d\n", value.BaseExperience)
	return nil
}
