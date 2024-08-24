package parser

import (
	"fmt"

	"govault/pkg/database"
	"govault/pkg/token"
)

func ParseInput(input string) {
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

	tokens := make([]token.Token, 0)

	for _, lexem := range lexems {
		switch lexem {
		case token.CREATE, token.DATABASE, token.TABLE, token.SELECT, token.FROM, token.WHERE, token.INSERT, token.INTO, token.VALUES:
			tokens = append(tokens, token.Token{Kind: token.KEYWORD, Value: lexem})
		case token.INT, token.STRING:
			tokens = append(tokens, token.Token{Kind: token.DATA_TYPE, Value: lexem})
		case token.EQUALS:
			tokens = append(tokens, token.Token{Kind: token.OPERATOR, Value: lexem})
		case token.OPEN_PARENTHESIS, token.CLOSE_PARENTHESIS, token.COMMA, token.SEMICOLON:
			tokens = append(tokens, token.Token{Kind: token.PUNCTUATION, Value: lexem})
		default:
			tokens = append(tokens, token.Token{Kind: token.IDENTIFIER, Value: lexem})
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
	case token.CREATE:
		switch tokens[1].Value {
		case token.DATABASE:
			parseCreateDatabaseStatement(tokens)
		case token.TABLE:
			parseCreateTableStatement(tokens)
		}
	case token.INSERT:
		switch tokens[1].Value {
		case token.INTO:
			parseInsertStatement(tokens)
		}
	}
}

func parseCreateDatabaseStatement(tokens []token.Token) {
	if len(tokens) < 3 || tokens[0].Value != token.CREATE || tokens[1].Value != token.DATABASE {
		fmt.Println("Error: Invalid CREATE DATABASE statement.")
		return
	}

	databaseName := tokens[2]

	if databaseName.Kind != token.IDENTIFIER {
		fmt.Println("Error: Invalid database name.")
		return
	}

	database.CreateDatabase(databaseName.Value)
}

func parseCreateTableStatement(tokens []token.Token) {
	// TODO: remove this check as it is already checked in the outer switch statement
	if len(tokens) < 4 || tokens[0].Value != token.CREATE || tokens[1].Value != token.TABLE {
		fmt.Println("Error: Invalid CREATE TABLE statement.")
		return
	}

	tableName := tokens[2]

	if tableName.Kind != token.IDENTIFIER {
		fmt.Println("Error: Invalid table name.")
		return
	}

	if tokens[3].Value != token.OPEN_PARENTHESIS {
		fmt.Println("Error: Invalid column definitions.")
		return
	}

	columns := make(map[string]string)
	columnIndex := 4

	for columnIndex < len(tokens) {
		if tokens[columnIndex].Value == token.CLOSE_PARENTHESIS {
			break
		}

		if tokens[columnIndex].Kind != token.IDENTIFIER {
			fmt.Printf("Error: Expected column name but got '%s'.\n", tokens[columnIndex].Value)
			return
		}

		columnName := tokens[columnIndex].Value
		columnIndex++

		if columnIndex >= len(tokens) || tokens[columnIndex].Kind != token.DATA_TYPE {
			fmt.Printf("Error: Expected data type for column '%s' but got '%s'.\n", columnName, tokens[columnIndex].Value)
			return
		}

		dataType := tokens[columnIndex].Value
		columnIndex++

		columns[columnName] = dataType

		if columnIndex < len(tokens) && tokens[columnIndex].Value == token.COMMA {
			columnIndex++
			continue
		} else if columnIndex < len(tokens) && tokens[columnIndex].Value == token.CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after column definition.")
			return
		}
	}

	if columnIndex >= len(tokens) || tokens[columnIndex].Value != token.CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of column definitions.")
		return
	}

	if columnIndex+1 >= len(tokens) || tokens[columnIndex+1].Value != token.SEMICOLON {
		fmt.Println("Error: Expected ';' at the end of the statement.")
		return
	}

	fmt.Println("Table name:", tableName.Value)
	fmt.Println("Columns:", columns)
	for columnName, columnType := range columns {
		fmt.Printf("Column: %s %s\n", columnName, columnType)
	}

	// TODO: create table with fields
	database.CreateTable("default", tableName.Value, columns)
}

func parseInsertStatement(tokens []token.Token) {
	// TODO: remove this check as it is already checked in the outer switch statement
	if len(tokens) < 4 || tokens[0].Value != token.INSERT || tokens[1].Value != token.INTO {
		fmt.Println("Error: Invalid INSERT INTO statement.")
	}

	tableName := tokens[2]

	if tableName.Kind != token.IDENTIFIER {
		fmt.Println("Error: Invalid table name.")
		return
	}

	if tokens[3].Value != token.OPEN_PARENTHESIS {
		fmt.Println("Error: Invalid column names.")
		return
	}

	columnNames := make([]string, 0)
	columnIndex := 4

	for columnIndex < len(tokens) {
		if tokens[columnIndex].Value == token.CLOSE_PARENTHESIS {
			break
		}

		if tokens[columnIndex].Kind != token.IDENTIFIER {
			fmt.Printf("Error: Expected column name but got '%s'.\n", tokens[columnIndex].Value)
			return
		}

		columnName := tokens[columnIndex].Value
		columnIndex++

		columnNames = append(columnNames, columnName)

		if columnIndex < len(tokens) && tokens[columnIndex].Value == token.COMMA {
			columnIndex++
			continue
		} else if columnIndex < len(tokens) && tokens[columnIndex].Value == token.CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after column name.")
			return
		}
	}

	if columnIndex >= len(tokens) || tokens[columnIndex].Value != token.CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of column names.")
		return
	}

	if columnIndex+1 >= len(tokens) || tokens[columnIndex+1].Value != token.VALUES {
		fmt.Println("Error: Expected VALUES keyword.")
		return
	}

	if columnIndex+2 >= len(tokens) || tokens[columnIndex+2].Value != token.OPEN_PARENTHESIS {
		fmt.Println("Error: Expected '(' before values.")
		return
	}

	values := make([]string, 0)
	valueIndex := columnIndex + 3

	for valueIndex < len(tokens) {
		if tokens[valueIndex].Value == token.CLOSE_PARENTHESIS {
			break
		}

		if tokens[valueIndex].Kind != token.IDENTIFIER {
			fmt.Printf("Error: Expected value but got '%s'.\n", tokens[valueIndex].Value)
			return
		}

		value := tokens[valueIndex].Value
		valueIndex++

		values = append(values, value)

		if valueIndex < len(tokens) && tokens[valueIndex].Value == token.COMMA {
			valueIndex++
			continue
		} else if valueIndex < len(tokens) && tokens[valueIndex].Value == token.CLOSE_PARENTHESIS {
			break
		} else {
			fmt.Println("Error: Expected ',' or ')' after value.")
			return
		}
	}

	if valueIndex >= len(tokens) || tokens[valueIndex].Value != token.CLOSE_PARENTHESIS {
		fmt.Println("Error: Expected ')' at the end of values.")
		return
	}

	if valueIndex+1 >= len(tokens) || tokens[valueIndex+1].Value != token.SEMICOLON {
		fmt.Println("Error: Expected ';' at the end of the statement.")
		return
	}

	fmt.Println("Table name:", tableName.Value)
	fmt.Println("Columns:", columnNames)
	fmt.Println("Values:", values)

	// TODO: insert data into table
	//table := database.Databases["default"].Tables[tableName.Value]
	//row := make(map[string]string)
	//for i, columnName := range columnNames {
	//	row[columnName] = values[i]
	//	table.Rows = append(table.Rows, row)
	//	database.Databases["default"].Tables[tableName.Value] = table
	//}
	database.InsertIntoTable("default", tableName.Value, columnNames, values)
}
