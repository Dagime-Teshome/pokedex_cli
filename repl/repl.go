package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Dagime-Teshome/pokedex_cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	previous string
	next     string
}

func (c *config) setPrev(s *string) {
	if s == nil {
		c.previous = "null"
		return
	}
	c.previous = *s
}
func (c *config) setNext(s *string) {
	if s == nil {
		c.next = "null"
		return
	}
	c.next = *s
}

var cache = pokecache.Newcache(5 * time.Second)

var commandsMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Show help text",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show list of location in poke-world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Navigates to the previous list of locations",
		callback:    commandMapB,
	},
}

type result struct {
	Name string
	Url  string
}

type locations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []result
}

func StartRepl() {
	configVar := config{
		previous: "",
		next:     "",
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}
func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.next != "" {
		url = c.next
	}
	locations, err := getMap(url)
	if err != nil {
		return err
	}
	c.setNext(locations.Next)
	c.setPrev(locations.Previous)
	if locations != nil {
		printLocations(*locations)
		return nil
	}
	return fmt.Errorf("no locations found")
}

func commandMapB(c *config) error {
	if c.previous == "null" {
		fmt.Println("You are on the fist page")
		return nil
	}
	locations, err := getMap(c.previous)
	c.setNext(locations.Next)
	c.setPrev(locations.Previous)
	if err != nil {
		return err
	}
	if locations != nil {
		printLocations(*locations)
		return nil
	}
	return fmt.Errorf("no locations found")
}

func executeCommand(command string, conf *config) {
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

func getMap(url string) (*locations, error) {
	var body []byte
	data, exists := cache.Get(url)
	if !exists {

		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("Error making request to :%s.got %s", url, err.Error())
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Status: %s when making Request to %s", res.Status, url)
		}
		if err != nil {
			return nil, err
		}
		cache.Add(url, body)
		for key, _ := range cache.Data {
			fmt.Println("key-----------", key)
		}
	} else {
		fmt.Println("found in cache", url)
		body = data
	}
	locations := locations{}
	json.Unmarshal(body, &locations)
	return &locations, nil
}

func printLocations(l locations) {
	for _, value := range l.Results {
		fmt.Println(value.Name)
	}
}
