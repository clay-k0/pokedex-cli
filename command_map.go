package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	res, err := cfg.pokeAPIClient.ListLocationAreas(cfg.nextLocationAreaURL)
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
