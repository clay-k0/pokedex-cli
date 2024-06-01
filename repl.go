package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/twistingmercury/go-figure"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func execREPL(cfg *config) {
	startupMsg := figure.NewFigure("Pokedex CLI", "", true)
	startupMsg.Print()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")

		fmt.Print("pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		validCommands := getCommands()

		command, ok := validCommands[commandName]
		if !ok {
			fmt.Println("unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "lists 20 pokemon world location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists the previous 20 pokemon world location areas",
			callback:    commandMapB,
		},
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
