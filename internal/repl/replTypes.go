package repl

import (
	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*shared.Config) error
}
