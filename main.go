package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// keywords
	CREATE   = "CREATE"
	DATABASE = "DATABASE"
	TABLE    = "TABLE"
	SELECT   = "SELECT"
	FROM     = "FROM"
	WHERE    = "WHERE"
	INSERT   = "INSERT"
	INTO     = "INTO"
	VALUES   = "VALUES"
	// data types
	INT    = "INT"
	STRING = "STRING"
	// operators
	EQUALS = "="
	// others
	OPEN_PARENTHESIS  = "("
	CLOSE_PARENTHESIS = ")"
	COMMA             = ","
	SEMICOLON         = ";"
)

type Token struct {
	Kind  string
	Value string
}

func parseInput(input string) {
	fmt.Println("Parsing:", input)

	lexems := make([]string, 0)
	lexem := ""

	for _, character := range input {
		if character == ' ' {
			if lexem != "" {
				fmt.Println("Lexem:", lexem)
				lexems = append(lexems, lexem)
				lexem = ""
			}

			continue
		}

		if character == '(' || character == ')' || character == ',' {
			fmt.Println("Lexem:", lexem)

			if lexem != "" {
				lexems = append(lexems, lexem)
				lexem = ""
			}

			lexems = append(lexems, string(character))

			continue
		}

		lexem += string(character)
	}

	fmt.Println("Lexems:", lexems)
	for _, lexem := range lexems {
		fmt.Printf("Lexem: '%s'\n", lexem)
	}

	tokens := make([]Token, 0)

	for _, lexem := range lexems {
		switch lexem {
		case CREATE:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case DATABASE:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case TABLE:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case SELECT:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case FROM:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case WHERE:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case INSERT:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case INTO:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case VALUES:
			tokens = append(tokens, Token{Kind: "KEYWORD", Value: lexem})
		case INT:
			tokens = append(tokens, Token{Kind: "DATA_TYPE", Value: lexem})
		case STRING:
			tokens = append(tokens, Token{Kind: "DATA_TYPE", Value: lexem})
		case EQUALS:
			tokens = append(tokens, Token{Kind: "OPERATOR", Value: lexem})
		case OPEN_PARENTHESIS:
			tokens = append(tokens, Token{Kind: "OTHER", Value: lexem})
		case CLOSE_PARENTHESIS:
			tokens = append(tokens, Token{Kind: "OTHER", Value: lexem})
		case COMMA:
			tokens = append(tokens, Token{Kind: "OTHER", Value: lexem})
		case SEMICOLON:
			tokens = append(tokens, Token{Kind: "OTHER", Value: lexem})
		default:
			tokens = append(tokens, Token{Kind: "OTHER", Value: lexem})
		}
	}

	fmt.Println("Tokens:", tokens)
	for _, token := range tokens {
		fmt.Println("Token:", token)
	}
}

func main() {
	fmt.Println("Starting GoDB...")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString(';')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "QUIT" {
			fmt.Println("Quiting GoDB...")
			return
		}

		fmt.Println("Input:", input)
		parseInput(input)
	}
}
