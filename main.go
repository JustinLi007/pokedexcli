package main

import (
	"log"
	"time"

	"github.com/JustinLi007/pokedexcli/internal/database"
	"github.com/JustinLi007/pokedexcli/internal/pokeapi"
)

func main() {
	pokedexDB, err := database.NewPokedexDB("database.json")
	if err != nil {
		log.Fatal("Failed to load database")
	}

	cfg := &config{
		pokeapiClient: pokeapi.NewClient(time.Second*5, time.Minute*5),
		pokedex:       pokedexDB,
	}

	startRepl(cfg)
}
