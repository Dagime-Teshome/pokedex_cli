package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Dagime-Teshome/pokedex_cli/internal/shared"
)

func getMap(url string, conf *shared.Config) (*shared.Locations, error) {
	var body []byte
	locationsCache, ok := conf.Cache.Get(url)
	if ok {
		locations := shared.Locations{}
		json.Unmarshal(locationsCache, &locations)
		return &locations, nil
	}
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

	conf.Cache.Add(url, body)
	locations := shared.Locations{}
	json.Unmarshal(body, &locations)
	return &locations, nil
}

func CommandMapB(c *shared.Config) error {
	if c.Previous == "null" {
		fmt.Println("You are on the fist page")
		return nil
	}
	locations, err := getMap(c.Previous, c)
	c.SetNext(locations.Next)
	c.SetPrev(locations.Previous)
	if err != nil {
		return err
	}
	if locations != nil {
		printLocations(*locations)
		return nil
	}
	return fmt.Errorf("no locations found")
}

func CommandMapF(c *shared.Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	}
	locations, err := getMap(url, c)
	if err != nil {
		return err
	}
	c.SetNext(locations.Next)
	c.SetPrev(locations.Previous)
	if locations != nil {
		printLocations(*locations)
		return nil
	}
	return fmt.Errorf("no locations found")
}

func printLocations(l shared.Locations) {
	for _, value := range l.Results {
		fmt.Println(value.Name)
	}
}
