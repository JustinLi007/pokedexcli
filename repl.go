package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JustinLi007/pokedexcli/internal/pokeapi"
	"github.com/JustinLi007/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	cache            pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

		userInput := scanner.Text()
		userInputParts := cleanInput(userInput)

		if len(userInputParts) == 0 {
			continue
		}

		commandName := userInputParts[0]

		command, ok := commands[commandName]
		if !ok {
			fmt.Fprintf(os.Stdout, "%v: command not found\n", userInput)
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Fprintln(os.Stdout, err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	parts := strings.Fields(lowercase)
	return parts
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Get the next page of location areas",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page location areas",
			callback:    commandMapBack,
		},
		"help": {
			name:        "help",
			description: "Display usage info",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}
