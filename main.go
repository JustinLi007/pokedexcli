package main

import (
	"time"

	"github.com/JustinLi007/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Second * 5),
	}

	startRepl(cfg)
}
