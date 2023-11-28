package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanTokens(t *testing.T) {
	str := `/+-*{}();,.=!<>!=>=<===//====*****+-*
		   /"abc"
		   123
		   *123.45
		   true
		   func
		   abc`

	expectTokens := []Token{
		NewToken(TokenTypeSlash, "/", nil, 1),
		NewToken(TokenTypePlus, "+", nil, 1),
		NewToken(TokenTypeMinus, "-", nil, 1),
		NewToken(TokenTypeStar, "*", nil, 1),
		NewToken(TokenTypeLBrace, "{", nil, 1),
		NewToken(TokenTypeRBrace, "}", nil, 1),
		NewToken(TokenTypeLParen, "(", nil, 1),
		NewToken(TokenTypeRParen, ")", nil, 1),
		NewToken(TokenTypeSemicolon, ";", nil, 1),
		NewToken(TokenTypeComma, ",", nil, 1),
		NewToken(TokenTypeDot, ".", nil, 1),
		NewToken(TokenTypeEqual, "=", nil, 1),
		NewToken(TokenTypeBang, "!", nil, 1),
		NewToken(TokenTypeLess, "<", nil, 1),
		NewToken(TokenTypeGreater, ">", nil, 1),
		NewToken(TokenTypeBangEqual, "!=", nil, 1),
		NewToken(TokenTypeGreaterEqual, ">=", nil, 1),
		NewToken(TokenTypeLessEqual, "<=", nil, 1),
		NewToken(TokenTypeEqualEqual, "==", nil, 1),
		NewToken(TokenTypeSlash, "/", nil, 2),
		NewToken(TokenTypeString, `"abc"`, "abc", 2),
		NewToken(TokenTypeNumber, "123", float64(123), 3),
		NewToken(TokenTypeStar, "*", nil, 4),
		NewToken(TokenTypeNumber, "123.45", float64(123.45), 4),
		NewToken(TokenTypeTrue, "true", nil, 5),
		NewToken(TokenTypeFunc, "func", nil, 6),
		NewToken(TokenTypeIdentifier, "abc", nil, 7),
		NewToken(TokenTypeEOF, "", nil, 7),
	}

	scanner := NewScanner(str)
	tokens := scanner.ScanTokens()
	assert.Equal(t, expectTokens, tokens)
}
