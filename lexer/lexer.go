package lexer

import (
	"wukong.com/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

