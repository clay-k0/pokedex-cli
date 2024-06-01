package main

import (
	"time"

	"github.com/clay-k0/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeAPIClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeAPIClient: pokeapi.NewClient(time.Hour),
	}
	execREPL(&cfg)
}
