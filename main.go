package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	KEYWORD     = "KEYWORD"
	DATA_TYPE   = "DATA_TYPE"
	OPERATOR    = "OPERATOR"
	PUNCTUATION = "PUNCTUATION"
	IDENTIFIER  = "IDENTIFIER"
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
	// punctuation
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

	// Lexical Analysis:

	lexems := make([]string, 0)
	lexem := ""

	for _, character := range input {
		if character == ' ' {
			if lexem != "" {
				lexems = append(lexems, lexem)
				lexem = ""
			}

			continue
		}

		if character == '(' || character == ')' || character == ',' || character == ';' {
			if lexem != "" {
				lexems = append(lexems, lexem)
				lexem = ""
			}

			lexems = append(lexems, string(character))

			continue
		}

		lexem += string(character)
	}

	if lexem != "" {
		lexems = append(lexems, lexem)
	}

	fmt.Println("Lexems:", lexems)
	for _, lexem := range lexems {
		fmt.Printf("Lexem: '%s'\n", lexem)
	}

	// Tokenization:

	tokens := make([]Token, 0)

	for _, lexem := range lexems {
		switch lexem {
		case CREATE, DATABASE, TABLE, SELECT, FROM, WHERE, INSERT, INTO, VALUES:
			tokens = append(tokens, Token{Kind: KEYWORD, Value: lexem})
		case INT, STRING:
			tokens = append(tokens, Token{Kind: DATA_TYPE, Value: lexem})
		case EQUALS:
			tokens = append(tokens, Token{Kind: OPERATOR, Value: lexem})
		case OPEN_PARENTHESIS, CLOSE_PARENTHESIS, COMMA, SEMICOLON:
			tokens = append(tokens, Token{Kind: PUNCTUATION, Value: lexem})
		default:
			tokens = append(tokens, Token{Kind: IDENTIFIER, Value: lexem})
		}
	}

	fmt.Println("Tokens:", tokens)
	for _, token := range tokens {
		fmt.Println("Token:", token)
	}

	// Syntax Analysis:

	if tokens[0].Value != CREATE || tokens[1].Value != TABLE {
		fmt.Println("Error: Invalid CREATE statement.")
		return
	}

	tableName := tokens[2]

	if tableName.Kind != IDENTIFIER {
		fmt.Println("Error: Invalid table name.")
		return
	}

	if tokens[3].Value != OPEN_PARENTHESIS {
		fmt.Println("Error: Invalid column definitions.")
		return
	}

	columns := make(map[string]string)

	for i := 4; i < len(tokens); i++ {
		if tokens[i].Value == CLOSE_PARENTHESIS {
			break
		}

		if tokens[i].Kind == PUNCTUATION {
			continue
		}

		columnName := tokens[i]
		fmt.Println("Column name:", columnName)
		if columnName.Kind != IDENTIFIER {
			fmt.Println("Error: Invalid column name.")
			return
		}

		columnType := tokens[i+1]
		if columnType.Kind != DATA_TYPE {
			fmt.Println("Error: Invalid column type.")
			return
		}

		columns[columnName.Value] = columnType.Value

		i++
	}

	fmt.Println("Table name:", tableName.Value)
	fmt.Println("Columns:", columns)
	for columnName, columnType := range columns {
		fmt.Printf("Column: %s %s\n", columnName, columnType)
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
