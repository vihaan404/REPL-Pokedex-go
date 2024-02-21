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

		fmt.Println("echoing: ", text)

	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)

	word := strings.Fields(lowered)

	return word
}
