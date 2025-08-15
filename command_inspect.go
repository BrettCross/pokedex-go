package main

import (
	"errors"
	"fmt"

	"github.com/brettcross/pokedex-go/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: Incorrect number of arguments to catch command")
	}
	pokeName := args[0]
	pokemon, exists := cfg.pokedex[pokeName]
	if exists {
		printPokemonData(pokemon)
		return nil
	}
	return errors.New("ERROR: Pokemon not found in Pokedex")
}

func printPokemonData(pokemon pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}
}