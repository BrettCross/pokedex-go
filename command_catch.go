package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("ERROR: Incorrect number of arguments to catch command")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	threshold := 40
	random := rand.Intn(pokemon.BaseExperience)
	if random >= threshold {
		// pokemon escaped
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	// pokemon caught
	fmt.Printf("%s was caught!\n", pokemon.Name)

	// check if pokedex already has entry for pokemon
	if _, exists := cfg.pokedex[pokemon.Name]; exists {
		fmt.Printf("you already have %s registered in your pokedex!", pokemon.Name)
		return nil 
	}
	
	// add pokemon to pokedex
	cfg.pokedex[pokemon.Name] = pokemon

	return nil 
}