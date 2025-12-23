package commands

import (
	"fmt"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func CommandHelp(c *shared.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}
