package main

import (
	"bufio"
	"fmt"
	"os"
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
