package main

import (
	"fmt"
	"log"
	"regexp"
)

const (
	TOKEN_TNAME = iota
	TOKEN_RNAME
	TOKEN_REGEX
	TOKEN_EQ
	TOKEN_BAR
	TOKEN_SCOLON
	TOKEN_WS
)

type LineCol struct {
	line int
	col  int
}

type Token struct {
	kind int
	text string
	line int
	col  int
	len  int
}

type FindFunc func([]byte) []byte

var find []FindFunc = []FindFunc{
	regexp.MustCompile("^[A-Z_]+").Find,
	regexp.MustCompile("^[a-z_]+").Find,
	regexp.MustCompile("^\".*\"").Find,
	regexp.MustCompile("^=").Find,
	regexp.MustCompile("^[|]").Find,
	regexp.MustCompile("^;").Find,
	regexp.MustCompile("^[ \t\n]+").Find,
}

func GenLineCol(buf []byte) []LineCol {
	buf_l := len(buf)
	col := 1
	line := 1
	var lc []LineCol

	for buf_i := 0; buf_i < buf_l; buf_i++ {
		lc = append(lc, LineCol{
			line: line,
			col:  col,
		})
		if buf[buf_i] == '\n' {
			col = 1
			line++
		} else {
			col++
		}
	}

	return lc
}

func Tokenize(buf []byte) []Token {
	lc := GenLineCol(buf)
	i := 0
	l := len(lc)

	find_l := len(find)
	suc := false

	var match []byte
	var tokens []Token

	for i < l {
		for find_i := 0; find_i < find_l; find_i++ {
			match = find[find_i](buf[i:])
			if match != nil {
				match_l := len(match)
				tokens = append(tokens, Token{
					kind: TOKEN_REGEX,
					text: string(match),
					line: lc[i].line,
					col:  lc[i].col,
					len:  match_l,
				})
				i += match_l
				suc = true
				break
			}
		}

		if !suc {
			log.Fatal("Failed to tokenize!")
		}
		suc = false
	}

	return tokens
}

func PrintlnTokens(tokens []Token) {
	fmt.Print("[")
	if len(tokens) > 0 {
		fmt.Printf("%s(%d:%d)", tokens[0].text, tokens[0].line, tokens[0].col)
	}
	for _, t := range tokens[1:] {
		fmt.Printf(", %s(%d:%d)", t.text, t.line, t.col)
	}
	fmt.Print("]\n")
}

func PrintlnSource(tokens []Token) {
	for _, t := range tokens[:] {
		fmt.Printf("%s", t.text)
	}
	fmt.Print("\n")
}
