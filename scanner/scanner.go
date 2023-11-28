package scanner

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int // start of current lexeme
	current int // current char of a lexeme
	line    int // current line
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isEOF() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, NewToken(TokenTypeEOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.source[s.current]
	s.current++
	switch c {
	case '(':
		s.addToken(TokenTypeLParen)
	case ')':
		s.addToken(TokenTypeRParen)
	case '{':
		s.addToken(TokenTypeLBrace)
	case '}':
		s.addToken(TokenTypeRBrace)
	case ',':
		s.addToken(TokenTypeComma)
	case '.':
		s.addToken(TokenTypeDot)
	case '-':
		s.addToken(TokenTypeMinus)
	case '+':
		s.addToken(TokenTypePlus)
	case ';':
		s.addToken(TokenTypeSemicolon)
	case '*':
		s.addToken(TokenTypeStar)
	case '!':
		if s.match('=') {
			s.addToken(TokenTypeBangEqual)
		} else {
			s.addToken(TokenTypeBang)
		}
	case '=':
		if s.match('=') {
			s.addToken(TokenTypeEqualEqual)
		} else {
			s.addToken(TokenTypeEqual)
		}
	case '>':
		if s.match('=') {
			s.addToken(TokenTypeGreaterEqual)
		} else {
			s.addToken(TokenTypeGreater)
		}

	case '<':
		if s.match('=') {
			s.addToken(TokenTypeLessEqual)
		} else {
			s.addToken(TokenTypeLess)
		}
	case '/':
		if s.match('/') {
			// These are comments
			for s.current < len(s.source) && s.source[s.current] != '\n' {
				s.current++
			}
		} else {
			s.addToken(TokenTypeSlash)
		}
	case ' ', '\t', '\r':
		// Ignore whitespace
	case '\n':
		s.line++
	case '"':
		s.scanString()
	default:
		if s.isDigit(c) {
			s.scanNumber()
		} else if s.isAlpha(c) {
			s.scanIdentifier()
		} else {
			fmt.Println("Unexpected character: "+string(c)+", line: ", s.line)
		}
	}
}

func (s *Scanner) scanString() {
	for s.source[s.current] != '"' && !s.isEOF() {
		if s.source[s.current] == '\n' {
			s.line++
		}
		s.current++
	}
	if s.isEOF() {
		fmt.Println("Unterminated string, line: ", s.line)
		return
	}
	// should be the end '"' now
	s.current++
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(TokenTypeString, value)
}

func (s *Scanner) scanNumber() {
	// Lox supports number format "123" or "123.45",
	// doesn't support ".123" or "123."
	consumeDigits := func() {
		for !s.isEOF() && s.isDigit(s.source[s.current]) {
			s.current++
		}
	}

	consumeDigits()
	if !s.isEOF() && s.source[s.current] == '.' && s.current+1 < len(s.source) && s.isDigit(s.source[s.current+1]) {
		s.current++
		consumeDigits()
	}
	value := s.source[s.start:s.current]
	f64, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return
	}
	s.addTokenLiteral(TokenTypeNumber, f64)
}

func (s *Scanner) scanIdentifier() {
	// For a identifier, the first character should be '_' or a letter
	// The rest should be a letter or digit
	s.current++
	for !s.isEOF() && (s.isAlpha(s.source[s.current]) || s.isDigit(s.source[s.current])) {
		s.current++
	}
	identifier := s.source[s.start:s.current]
	tokenType, exist := keywords[identifier]
	if !exist {
		tokenType = TokenTypeIdentifier
	}
	s.addToken(tokenType)
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.tokens = append(s.tokens, NewToken(tokenType, s.source[s.start:s.current], nil, s.line))
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, NewToken(tokenType, s.source[s.start:s.current], literal, s.line))
}

func (s *Scanner) match(expected byte) bool {
	if s.isEOF() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) isEOF() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}
