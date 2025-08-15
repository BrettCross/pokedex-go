package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	// get area to explore from args
	if len(args) == 0 {
		return errors.New("too few arguments")
	}
	if len(args) > 1 {
		fmt.Printf("WARNING: Expecting 1 argument. Received %d\n", len(args))
	}
	exploreResp, err := cfg.pokeapiClient.ExploreArea(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}