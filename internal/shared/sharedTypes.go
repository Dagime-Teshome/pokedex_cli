package shared

import (
	"github.com/Dagime-Teshome/pokedex_cli/internal/pokecache"
)

type Config struct {
	Previous string
	Next     string
	Cache    pokecache.Cache
}

func (c *Config) SetPrev(s *string) {
	if s == nil {
		c.Previous = "null"
		return
	}
	c.Previous = *s
}
func (c *Config) SetNext(s *string) {
	if s == nil {
		c.Next = "null"
		return
	}
	c.Next = *s
}

type Locations struct {
	Count    int
	Next     *string
	Previous *string
	Results  []result
}

type result struct {
	Name string
	Url  string
}
