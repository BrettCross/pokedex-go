package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/brettcross/pokedex-go/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		value, ok := getCommands()[commandName]
		if ok {
			err := value.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand {
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display names of next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display names of previous 20 locations",
			callback:    commandMapb,
		},
	}
}
