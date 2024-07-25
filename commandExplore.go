package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Must provide a location name")
	}

	locationName := args[0]
	fmt.Printf("Exploring %v...\n", locationName)
	locationDetail, err := cfg.pokeapiClient.GetLocationDetails(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, value := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %v\n", value.Pokemon.Name)
	}

	return nil
}
