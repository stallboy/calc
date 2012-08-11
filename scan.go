package main

import (
	"container/list"
	"unicode"
)

type TokenType uint32

const (
	PLUS TokenType = iota
	MINUS
	MUL
	DIV
	POWER
	ASSIGN
	LPAREN
	RPAREN
	ID
	NUM
	EOL //end of line
)

type Token struct {
	Type TokenType
	Lit  string
}

type Scanner struct {
	cmd     []rune
	pos     int
	untoken list.List
}

func NewScanner(cmd string) *Scanner {
	r := new(Scanner)
	r.cmd = []rune(cmd)
	return r
}

func tok(tp TokenType) Token {
	return Token{Type: tp}
}

func (s *Scanner) Scan() Token {
	if e := s.untoken.Front(); e != nil {
		s.untoken.Remove(e)
		return e.Value.(Token)
	}

	var ch rune
	cmdlen := len(s.cmd)
	for s.pos < cmdlen {
		ch = s.cmd[s.pos]
		if ch != ' ' && ch != '\t' {
			break
		}

		s.pos++
	}

	if s.pos == cmdlen {
		return tok(EOL)
	}

	switch ch {
	case '+':
		s.pos++
		return tok(PLUS)
	case '-':
		s.pos++
		return tok(MINUS)
	case '*':
		s.pos++
		if s.pos < cmdlen {
			ch1 := s.cmd[s.pos]
			if ch1 == '*' {
				s.pos++
				return tok(POWER)
			}
		}
		return tok(MUL)
	case '/':
		s.pos++
		return tok(DIV)
	case '=':
		s.pos++
		return tok(ASSIGN)
	case '(':
		s.pos++
		return tok(LPAREN)
	case ')':
		s.pos++
		return tok(RPAREN)
	default:
		if unicode.IsDigit(ch) {
			num := string(ch)

			state := 0
			end := false
			for !end {
				s.pos++
				if s.pos == cmdlen {
					break
				}
				ch = s.cmd[s.pos]

				switch state {
				case 0:
					if unicode.IsDigit(ch) || ch == '.' {
						num += string(ch)
					} else if ch == 'e' {
						num += string(ch)
						state = 1
					} else {
						end = true
					}
				case 1:
					if unicode.IsDigit(ch) {
						num += string(ch)
						state = 2
					} else if ch == '+' || ch == '-' {
						num += string(ch)
					} else {
						panic("scan num format err " + num)
						end = true
					}
				case 2:
					if unicode.IsDigit(ch) {
						num += string(ch)
					} else {
						end = true
					}
				}
			}
			return Token{NUM, num}

		} else if unicode.IsLetter(ch) {
			id := string(ch)
			for {
				s.pos++
				if s.pos == cmdlen {
					break
				}
				ch = s.cmd[s.pos]

				if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
					id += string(ch)
				} else {
					break
				}
			}
			return Token{ID, id}
		}

	}

	panic("unknown rune " + string(ch))

}

func (s *Scanner) Unscan(token Token) {
	s.untoken.PushFront(token)
}
