package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JustinLi007/pokedexcli/internal/database"
	"github.com/JustinLi007/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          *database.PokedexDB
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
				return
			}
		}

		userInput := scanner.Text()
		userInputParts := cleanInput(userInput)

		if len(userInputParts) == 0 {
			continue
		}

		commandName := userInputParts[0]
		args := make([]string, 0)
		if len(userInputParts) > 1 {
			args = userInputParts[1:]
		}

		command, ok := commands[commandName]
		if !ok {
			fmt.Fprintf(os.Stdout, "%v: command not found\n", userInput)
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
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
		"explore": {
			name:        "explore",
			description: "Display information about a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Display information about a caught Pokemon",
			callback:    commandInspect,
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
