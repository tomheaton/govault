package lexer

import (
	"fmt"
	"govault/pkg/token"
)

func Tokenize(input string) []token.Token {
	fmt.Println("Parsing:", input)

	// Lexical Analysis
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

	// Tokenization
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

	return tokens
}
