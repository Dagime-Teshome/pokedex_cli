package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Dagime-Teshome/pokedex_cli/internal/commands"
	"github.com/Dagime-Teshome/pokedex_cli/internal/pokecache"
	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

var commandsMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.CommandExit,
	},
	"help": {
		name:        "help",
		description: "Show help text",
		callback:    commands.CommandHelp,
	},
	"map": {
		name:        "map",
		description: "Show list of location in poke-world",
		callback:    commands.CommandMapF,
	},
	"mapb": {
		name:        "mapb",
		description: "Navigates to the previous list of locations",
		callback:    commands.CommandMapB,
	},
}

func StartRepl() {
	configVar := shared.Config{
		Previous: "",
		Next:     "",
		Cache:    *pokecache.Newcache(10 * time.Second),
	}
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		executeCommand(commandName, &configVar)
	}
}
func executeCommand(command string, conf *shared.Config) {
	_, exists := commandsMap[command]
	if exists {
		err := commandsMap[command].callback(conf)
		if err != nil {
			fmt.Println("Error executing command: ", err)
		}
		return
	}
	fmt.Println("Unknown Command")
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
