package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"govault/pkg/database"
	"govault/pkg/parser"
)

func main() {
	fmt.Println("Starting GoVault...")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString(';')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "QUIT;" {
			fmt.Println("Quiting GoVault...")
			break
			//return
		}

		fmt.Println("Input:", input)
		parser.ParseInput(input)
	}

	fmt.Println(database.Databases)
}
