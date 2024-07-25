package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	caughtPokemon, err := cfg.pokedex.GetPokemons()
	if err != nil {
		return err
	}

	fmt.Println("Your Pokedex:")
	for _, value := range caughtPokemon {
		fmt.Printf(" - %v\n", value.Name)
	}

	return nil
}
