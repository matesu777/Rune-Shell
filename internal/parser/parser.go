package parser

type Parser struct {
	lexer Lexer
	curr  Token
}

func NewParser(l Lexer) *Parser {

	p := &Parser{
		lexer: l,
	}

	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curr = p.lexer.NextToken()
}

func (p *Parser) ParseCommand() Command {

	cmd := Command{}

	// primeiro token precisa ser o comando

	if p.curr.Type != COMMAND {
		return cmd
	}

	cmd.Name = p.curr.Value

	p.nextToken()

	// resto vira argumento

	for p.curr.Type == STRING ||
		p.curr.Type == COMMAND {

		cmd.Args = append(
			cmd.Args,
			p.curr.Value,
		)

		p.nextToken()
	}

	return cmd
}
