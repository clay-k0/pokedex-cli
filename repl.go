package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
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
			fmt.Printf("unknown command '%v'\n", commandName)
			continue
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
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
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
