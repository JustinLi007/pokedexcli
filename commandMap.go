package main

import (
	"errors"
	"fmt"
)

func commandMapForward(cfg *config, arts ...string) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.prevLocationsURL = response.Previous

	for _, v := range response.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("Top of page reached")
	}

	response, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.prevLocationsURL = response.Previous

	for _, v := range response.Results {
		fmt.Println(v.Name)
	}

	return nil
}
