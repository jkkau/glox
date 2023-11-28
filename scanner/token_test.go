package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	token := NewToken(TokenTypeEOF, "", nil, 1)
	assert.Equal(t, "Type: EOF, Lexeme: , Line: 1", token.String())
}
