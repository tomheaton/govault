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

func parseCreateDatabaseStatement(tokens []Token) {
	if len(tokens) < 3 || tokens[0].Value != CREATE || tokens[1].Value != DATABASE {
		fmt.Println("Error: Invalid CREATE DATABASE statement.")
		return
	}

	databaseName := tokens[2]

	if databaseName.Kind != IDENTIFIER {
		fmt.Println("Error: Invalid database name.")
		return
	}

	fmt.Println("Database name:", databaseName.Value)
}

func parseCreateTableStatement(tokens []Token) {
	if len(tokens) < 4 || tokens[0].Value != CREATE || tokens[1].Value != TABLE {
		fmt.Println("Error: Invalid CREATE TABLE statement.")
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
	columnIndex := 4

	for columnIndex < len(tokens) {
		if tokens[columnIndex].Value == CLOSE_PARENTHESIS {
			break
		}

		if tokens[columnIndex].Kind != IDENTIFIER {
			fmt.Printf("Error: Expected column name but got '%s'.\n", tokens[columnIndex].Value)
			return
		}

		columnName := tokens[columnIndex].Value
		columnIndex++

		if columnIndex >= len(tokens) || tokens[columnIndex].Kind != DATA_TYPE {
			fmt.Printf("Error: Expected data type for column '%s' but got '%s'.\n", columnName, tokens[columnIndex].Value)
			return
		}

		dataType := tokens[columnIndex].Value
		columnIndex++

		columns[columnName] = dataType

		if columnIndex < len(tokens) && tokens[columnIndex].Value == COMMA {
			columnIndex++
			continue
		} else if columnIndex < len(tokens) && tokens[columnIndex].Value == CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after column definition.")
			return
		}
	}

	if columnIndex >= len(tokens) || tokens[columnIndex].Value != CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of column definitions.")
		return
	}

	if columnIndex+1 >= len(tokens) || tokens[columnIndex+1].Value != SEMICOLON {
		fmt.Println("Error: Expected ';' at the end of the statement.")
		return
	}

	fmt.Println("Table name:", tableName.Value)
	fmt.Println("Columns:", columns)
	for columnName, columnType := range columns {
		fmt.Printf("Column: %s %s\n", columnName, columnType)
	}

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

	if len(tokens) < 3 {
		fmt.Println("Error: Invalid statement.")
		return
	}

	switch tokens[0].Value {
	case CREATE:
		switch tokens[1].Value {
		case DATABASE:
			parseCreateDatabaseStatement(tokens)
		case TABLE:
			parseCreateTableStatement(tokens)
		}
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
