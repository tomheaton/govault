package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Table struct {
	Name    string
	Columns map[string]string
	Rows    []map[string]string
}

type Database struct {
	Name   string
	Tables map[string]Table
}

var databases = make(map[string]Database)

// TODO: prepend with TOKEN_
const (
	KEYWORD     = "KEYWORD"
	DATA_TYPE   = "DATA_TYPE"
	OPERATOR    = "OPERATOR"
	PUNCTUATION = "PUNCTUATION"
	IDENTIFIER  = "IDENTIFIER"
)

// TODO: prepend with KEYWORD_
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
	// TODO: remove this check as it is already checked in the outer switch statement
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

	// TODO: create database
	databases[databaseName.Value] = Database{Name: databaseName.Value, Tables: make(map[string]Table)}
}

func parseCreateTableStatement(tokens []Token) {
	// TODO: remove this check as it is already checked in the outer switch statement
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

	// TODO: create table with fields
	databases["default"].Tables[tableName.Value] = Table{Name: tableName.Value, Columns: columns, Rows: make([]map[string]string, 0)}
}

func parseInsertStatement(tokens []Token) {
	// TODO: remove this check as it is already checked in the outer switch statement
	if len(tokens) < 4 || tokens[0].Value != INSERT || tokens[1].Value != INTO {
		fmt.Println("Error: Invalid INSERT INTO statement.")
	}

	tableName := tokens[2]

	if tableName.Kind != IDENTIFIER {
		fmt.Println("Error: Invalid table name.")
		return
	}

	if tokens[3].Value != OPEN_PARENTHESIS {
		fmt.Println("Error: Invalid column names.")
		return
	}

	columnNames := make([]string, 0)
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

		columnNames = append(columnNames, columnName)

		if columnIndex < len(tokens) && tokens[columnIndex].Value == COMMA {
			columnIndex++
			continue
		} else if columnIndex < len(tokens) && tokens[columnIndex].Value == CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after column name.")
			return
		}
	}

	if columnIndex >= len(tokens) || tokens[columnIndex].Value != CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of column names.")
		return
	}

	if columnIndex+1 >= len(tokens) || tokens[columnIndex+1].Value != VALUES {
		fmt.Println("Error: Expected VALUES keyword.")
		return
	}

	if columnIndex+2 >= len(tokens) || tokens[columnIndex+2].Value != OPEN_PARENTHESIS {
		fmt.Println("Error: Expected '(' before values.")
		return
	}

	values := make([]string, 0)
	valueIndex := columnIndex + 3

	for valueIndex < len(tokens) {
		if tokens[valueIndex].Value == CLOSE_PARENTHESIS {
			break
		}

		if tokens[valueIndex].Kind != IDENTIFIER {
			fmt.Printf("Error: Expected value but got '%s'.\n", tokens[valueIndex].Value)
			return
		}

		value := tokens[valueIndex].Value
		valueIndex++

		values = append(values, value)

		if valueIndex < len(tokens) && tokens[valueIndex].Value == COMMA {
			valueIndex++
			continue
		} else if valueIndex < len(tokens) && tokens[valueIndex].Value == CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after value.")
			return
		}
	}

	if valueIndex >= len(tokens) || tokens[valueIndex].Value != CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of values.")
		return
	}

	if valueIndex+1 >= len(tokens) || tokens[valueIndex+1].Value != SEMICOLON {
		fmt.Println("Error: Expected ';' at the end of the statement.")
		return
	}

	fmt.Println("Table name:", tableName.Value)
	fmt.Println("Columns:", columnNames)
	fmt.Println("Values:", values)

	// TODO: insert data into table
	table := databases["default"].Tables[tableName.Value]
	row := make(map[string]string)
	for i, columnName := range columnNames {
		row[columnName] = values[i]
		table.Rows = append(table.Rows, row)
		databases["default"].Tables[tableName.Value] = table
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
	case INSERT:
		switch tokens[1].Value {
		case INTO:
			parseInsertStatement(tokens)
		}
	}
}

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
		parseInput(input)
	}

	fmt.Println(databases)
}
