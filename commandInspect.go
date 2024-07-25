package main

import (
	"errors"
	"fmt"

	"github.com/JustinLi007/pokedexcli/internal/database"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide Pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokedex.GetPokemon(pokemonName)
	if err != nil {
		if errors.Is(err, database.ErrNotExist) {
			return errors.New("You have not caught that Pokemon")
		}
		return err
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, value := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemon.Types {
		fmt.Printf(" - %v\n", value.Type.Name)
	}

	return nil
}
