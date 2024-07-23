package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("\nerror: you cannot go to the previous page from first the page")
	}
	res, err := cfg.pokeAPIClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("---LOCATION AREAS---")
	for _, area := range res.Results {
		fmt.Printf("> %s\n", area.Name)
	}
	fmt.Println("")

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}
