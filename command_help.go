package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex")
	availableCommand := commandCalling()

	for _, cmd := range availableCommand {
		fmt.Printf(" - %s -  %s \n", cmd.name, cmd.discription)
	}
	return nil
}
