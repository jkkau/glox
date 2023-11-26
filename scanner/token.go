package scanner

type TokenType int

type Token[T any] struct {
	Type    TokenType
	Lexeme  string
	Literal T
	Line    int
}

func NewToken[T any](tokenType TokenType, lexeme string, literal T, line int) Token[T] {
	return Token[T]{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}