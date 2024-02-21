package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	fmt.Print(">")

	scaner.Scan()

	text := scaner.Text()

	fmt.Println("echoing: ", text)

}
