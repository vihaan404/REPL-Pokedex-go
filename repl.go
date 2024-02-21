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

		command := cleanText[0]
		switch command {
		case "exit":
			os.Exit(0)

		}

		if len(cleanText) == 0 {
			fmt.Println()
			continue
		}

		fmt.Println("echoing: ", cleanText)

	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	word := strings.Fields(lowered)

	return word
}
