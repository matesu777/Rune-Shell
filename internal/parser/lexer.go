package parser

import (
	"unicode"
)

type Lexer struct {
	buf    string
	cursor int
	ch     byte
}

func NewLexer(buf string) Lexer {
	return Lexer{
		buf:    buf,
		cursor: 0,
		ch:     0,
	}
}

func (l *Lexer) IsEOF() bool {
	return l.cursor >= len(l.buf)
}

func (l *Lexer) Peek() byte {
	if l.IsEOF() {
		return 0
	}
	return l.buf[l.cursor]
}

func (l *Lexer) Next() byte {
	if l.IsEOF() {
		return 0
	}
	ch := l.buf[l.cursor]
	l.cursor++
	return ch
}

func (l *Lexer) Expect(exp byte) bool {
	if l.IsEOF() {
		return false
	}
	return l.buf[l.cursor] == exp
}

func (l *Lexer) ConsumeWhiteSpaces() {
	for !l.IsEOF() && unicode.IsSpace(rune(l.Peek())) {
		l.cursor++
	}
	// fmt.Println("Espaço consumido")
}

func (l *Lexer) NextToken() Token {
	l.ConsumeWhiteSpaces()

	if l.IsEOF() {
		return NewToken(EOF, "", SingleQuote)
	}

	char := l.Peek()

	switch char {
	case '\'', '"':
		value, quote := l.ParserString()
		return NewToken(STRING, value, quote)
	default:
		value, quote := l.ParserCommand()
		return NewToken(COMMAND, value, quote)
	}
}

func (l *Lexer) ParserString() (string, QuoteType) {

	quote := l.Peek()

	var quoteType QuoteType

	if quote == '\'' {
		quoteType = SingleQuote
	} else {
		quoteType = DoubleQuote
	}

	l.cursor++

	var str []byte

	for !l.IsEOF() && l.Peek() != quote {
		str = append(str, l.Peek())
		l.cursor++
	}

	if l.IsEOF() {
		panic("string não fechada")
	}

	l.cursor++

	return string(str), quoteType
}

func (l *Lexer) ParserCommand() (string, QuoteType) {
	var buf []byte

	for !l.IsEOF() {
		ch := l.Peek()

		if unicode.IsSpace(rune(ch)) {
			break
		}

		buf = append(buf, ch)
		l.cursor++
	}

	return string(buf), SingleQuote
}
