package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
<addop>	-> + | -
<mulop>	-> * | /
<powop>	-> **

<command>	-> [id = ] <exp> eol
<exp>		-> <mulexp> { <addop> <mulexp> }
<mulexp>	-> <powexp> { <mulop> <powexp> }
<powexp>	-> <unaryexp> [ <powop> <unaryexp> ]
<unaryexp>	-> [<addop>] <atom>
<atom>		-> (<exp>) | num | id
*/

var ctx = map[string]float64{
	"e":  2.7182818,
	"pi": 3.1415926,
}

var scanner *Scanner
var cur Token

func scan() {
	cur = scanner.Scan()
}

func match(tp TokenType) {
	if cur.Type == tp {
		scan()
		return
	}

	panic(fmt.Sprintf("unmatched token %v, expected %v", cur, tp))
}

func Parse(cmd string) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("panic: %v\n", x)
		}
	}()

	scanner = NewScanner(cmd)
	f := scanner.Scan()
	s := scanner.Scan()

	if f.Type == ID && s.Type == ASSIGN {
		scan()
		r := exp()
		match(EOL)
		ctx[f.Lit] = r
		fmt.Printf("%s=%v\n", f.Lit, r)

	} else {
		scanner.Unscan(s)
		scanner.Unscan(f)
		scan()
		r := exp()
		match(EOL)
		fmt.Println(r)
	}
}

func exp() float64 {
	r := multiexp()

	var tp TokenType
	for tp = cur.Type; tp == PLUS || tp == MINUS; tp = cur.Type {
		scan()
		r2 := multiexp()

		if tp == PLUS {
			r += r2
		} else {
			r -= r2
		}
	}

	return r
}

func multiexp() float64 {
	r := powexp()

	var tp TokenType
	for tp = cur.Type; tp == MUL || tp == DIV; tp = cur.Type {
		scan()
		r2 := powexp()

		if tp == MUL {
			r *= r2
		} else {
			r /= r2
		}
	}
	return r
}

func powexp() float64 {
	r := unaryexp()
	for cur.Type == POWER {
		scan()
		r2 := unaryexp()
		r = math.Pow(r, r2)
	}
	return r
}

func unaryexp() float64 {
	switch cur.Type {
	case PLUS:
		scan()
		return atom()
	case MINUS:
		scan()
		return -atom()
	}
	return atom()
}

func atom() float64 {
	switch cur.Type {
	case LPAREN:
		match(LPAREN)
		r := exp()
		match(RPAREN)
		return r
	case ID:
		lit := cur.Lit
		scan()
		if v, ok := ctx[lit]; ok {
			return v
		}
		panic("unknown id " + lit)
	case NUM:
		lit := cur.Lit
		scan()
		if f, e := strconv.ParseFloat(lit, 64); e == nil {
			return f
		}
		panic("parse float error " + lit)
	}
	panic(fmt.Sprintf("unexpected token %v", cur))
}
