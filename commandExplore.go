package main

import "fmt"

func commandExplore(cfg *config, locationName string) error {
	fmt.Printf("Exploring %v...\n", locationName)
	encounters, err := cfg.pokeapiClient.ListLocationDetails(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, value := range encounters.PokemonEncounters {
		fmt.Printf(" - %v\n", value.Pokemon.Name)
	}

	return nil
}
