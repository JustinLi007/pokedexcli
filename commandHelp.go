package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println()

	return nil
}
