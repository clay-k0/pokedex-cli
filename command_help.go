package main

import "fmt"

func commandHelp() {
	fmt.Println("")
	fmt.Println("---POKEDEX HELP---")
	fmt.Println("available commands:")

	validCommands := getCommands()
	for _, cmd := range validCommands {
		fmt.Printf("> %s - %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")
}
