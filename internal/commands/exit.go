package commands

import (
	"fmt"
	"os"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func CommandExit(c *shared.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
