package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processCommand(command string) {
	fmt.Println("Processing:", command)

	parts := strings.Split(command, " ")

	if len(parts) == 0 {
		fmt.Println("Invalid command")
		return
	}

	switch strings.ToLower(parts[0]) {
	case "create":
		switch strings.ToLower(parts[1]) {
		case "table":
			tableName := parts[2]
			fmt.Println("Creating table:", tableName)
		}
	}
}

func main() {
	fmt.Println("Starting GoDB...")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "quit" || input == "QUIT" {
			fmt.Println("Quiting GoDB...")
			return
		}

		fmt.Println("Input:", input)
		processCommand(input)
	}
}
