package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {

	scaner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scaner.Scan()

		text := scaner.Text()

		cleanText := cleanInput(text)
		if len(cleanText) == 0 {
			fmt.Println()
			continue
		}

		commandName := cleanText[0]

		availableCommand := commandCalling()

		command, ok := availableCommand[commandName]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.callback()

	}

}

type cliCommand struct {
	name        string
	discription string
	callback    func() error
}

func commandCalling() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			discription: "Its a help menu what you dont understand",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			discription: "For exiting the pokedex",
			callback:    commandExit,
		},
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	word := strings.Fields(lowered)

	return word
}
