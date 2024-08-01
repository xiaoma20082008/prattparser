//
// File: parser.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:11:22
// Last Modified By: xiaoma20082008 (mmccxx2519@gmail.com>)
// ------------------------------------------------------------------------
//
// Copyright (C) xiaoma20082008. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package prattparser

import (
	"bytes"
	"fmt"
)

type Parser struct {
	S Scanner
	L Lexed
}

func NewParser(src string) *Parser {
	return &Parser{S: Scanner{src: bytes.NewBuffer([]byte(src))}}
}

func (p *Parser) Parse() Expr {
	p.L = p.S.Scan()
	return p.expr(LowestPrec)
}

func (p *Parser) expr(rbp int) Expr {
	t := p.L
	p.L = p.S.Scan()

	var left Expr
	switch t.Tok {
	case ILLEGAL:
		panic(fmt.Sprintf("Parse Error: %s", t.Lit))
	case LP:
		left = p.expr(0)
		if p.L.Tok != RP {
			panic("Parse Error: expected )")
		}
		p.L = p.S.Scan()
	default:
		switch true {
		case t.Tok.IsScalar():
			left = ScalarExpr{Val: t.Lit, Typ: t.Tok}
		case t.Tok.IsUnary():
			left = UnaryExpr{X: p.expr(HighestPrec), Op: t.Tok}
		}
	}

	// left binding power
	for rbp < p.L.Tok.Precedence() {
		t = p.L
		p.L = p.S.Scan()
		//left denotation
		switch true {
		case t.Tok.IsRightAssoc():
			left = BinaryExpr{L: left, Op: t.Tok, R: p.expr(t.Tok.Precedence() - 1)}
		case t.Tok.IsLeftAssoc():
			left = BinaryExpr{L: left, Op: t.Tok, R: p.expr(t.Tok.Precedence())}
		}
	}
	return left
}
