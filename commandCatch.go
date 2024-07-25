package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide Pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	// chance := rand.Float64() * float64(pokemon.BaseExperience)
	chance := rand.Intn(pokemon.BaseExperience)
	if chance > 40.0 {
		fmt.Printf("%v escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemonName)

	return cfg.pokedex.AddPokemon(pokemonName, pokemon)
}
