package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("ERROR: too many arguments")
	}

	if len(cfg.pokedex) == 0 {
		fmt.Println("You have no pokemon!")
		return nil
	}

	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}