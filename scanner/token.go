package scanner

import "fmt"

type TokenType int

const (
	TokenTypeEOF          TokenType = iota
	TokenTypePlus                   // +
	TokenTypeMinus                  // -
	TokenTypeStar                   // *
	TokenTypeSlash                  // /
	TokenTypeLBrace                 // {
	TokenTypeRBrace                 // }
	TokenTypeLParen                 // (
	TokenTypeRParen                 // )
	TokenTypeSemicolon              // ;
	TokenTypeComma                  // ,
	TokenTypeDot                    // .
	TokenTypeBang                   // !
	TokenTypeBangEqual              // !=
	TokenTypeLess                   // <
	TokenTypeLessEqual              // <=
	TokenTypeGreater                // >
	TokenTypeGreaterEqual           // >=
	TokenTypeEqual                  // =
	TokenTypeEqualEqual             // ==
	TokenTypeString                 // "xxx"
	TokenTypeNumber                 // 123
	TokenTypeIdentifier             // identifier
	TokenTypeIf                     // if
	TokenTypeElse                   // else
	TokenTypeWhile                  // while
	TokenTypeFor                    // for
	TokenTypeReturn                 // return
	TokenTypeAnd                    // and
	TokenTypeOr                     // or
	TokenTypeTrue                   // true
	TokenTypeFalse                  // false
	TokenTypeVar                    // var
	TokenTypeNil                    // nil
	TokenTypePrint                  // print
	TokenTypeClass                  // class
	TokenTypeThis                   // this
	TokenTypeSuper                  // super
	TokenTypeFunc                   // func
)

var tokenMap = map[TokenType]string{
	TokenTypeEOF:          "EOF",
	TokenTypePlus:         "PLUS",
	TokenTypeMinus:        "MINUS",
	TokenTypeStar:         "STAR",
	TokenTypeSlash:        "SLASH",
	TokenTypeLBrace:       "LBRACE",
	TokenTypeRBrace:       "RBRACE",
	TokenTypeLParen:       "LPAREN",
	TokenTypeRParen:       "RPAREN",
	TokenTypeSemicolon:    "SEMICOLON",
	TokenTypeComma:        "COMMA",
	TokenTypeDot:          "DOT",
	TokenTypeBang:         "BANG",
	TokenTypeBangEqual:    "BANGEQUAL",
	TokenTypeLess:         "LESS",
	TokenTypeLessEqual:    "LESSEQUAL",
	TokenTypeGreater:      "GREATER",
	TokenTypeGreaterEqual: "GREATEREQUAL",
	TokenTypeEqual:        "EQUAL",
	TokenTypeEqualEqual:   "EQUALEQUAL",
	TokenTypeString:       "STRING",
	TokenTypeNumber:       "NUMBER",
	TokenTypeIdentifier:   "IDENTIFIER",
}

var keywords = map[string]TokenType{
	"if":     TokenTypeIf,
	"else":   TokenTypeElse,
	"while":  TokenTypeWhile,
	"for":    TokenTypeFor,
	"return": TokenTypeReturn,
	"and":    TokenTypeAnd,
	"or":     TokenTypeOr,
	"true":   TokenTypeTrue,
	"false":  TokenTypeFalse,
	"var":    TokenTypeVar,
	"nil":    TokenTypeNil,
	"print":  TokenTypePrint,
	"class":  TokenTypeClass,
	"this":   TokenTypeThis,
	"super":  TokenTypeSuper,
	"func":   TokenTypeFunc,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprint("Type: ", tokenMap[t.Type], ", Lexeme: ", t.Lexeme, ", Line: ", t.Line)
}
