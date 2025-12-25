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
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Println("Stats: ")
	for _, value := range value.Stats {
		fmt.Printf("  - %s : %d\n", value.Stat.Name, value.Base_Stat)
	}
	fmt.Println("Types: ")
	for _, value := range value.Types {
		fmt.Printf("  - %s\n", value.Type.Name)
	}
	return nil
}
