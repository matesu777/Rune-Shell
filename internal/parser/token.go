package parser

type TokenType int
type QuoteType int

const (
	NoQuote QuoteType = iota
	SingleQuote
	DoubleQuote
)

const (
	EOF TokenType = iota
	STRING
	COMMAND
	FLAG
)

type Token struct {
	Type  TokenType
	Value string
	Quote QuoteType
}

var tokenNames = []string{
	EOF:     "EOF",
	STRING:  "STRING",
	COMMAND: "COMMAND",
	FLAG:    "FLAG",
}

func (t TokenType) String() string {
	if t < 0 || int(t) >= len(tokenNames) {
		return "UNKNOWN"
	}
	return tokenNames[t]
}

func NewToken(t TokenType, v string, q QuoteType) Token {
	return Token{
		Type:  t,
		Value: v,
		Quote: q,
	}
}
