package token

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
