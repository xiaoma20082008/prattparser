//
// File: token.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:11:03
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

type Token int

const (
	ILLEGAL Token = iota
	EOF
	INT // int
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %
	POW // ^
	LP  // (
	RP  // )
)

var tokens = [...]string{
	ILLEGAL: "illegal",
	EOF:     "eof",
	INT:     "int",
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	QUO:     "/",
	REM:     "%",
	POW:     "^",
	LP:      "(",
	RP:      ")",
}

const (
	LowestPrec  = 0
	HighestPrec = 100
)

func (tok Token) Precedence() int {
	prec := LowestPrec
	switch tok {
	case REM: // %
		prec = 7
	case POW: // ^
		prec = 6
	case MUL, QUO: // *,/
		prec = 5
	case ADD, SUB: // +,-
		prec = 4
	}
	return prec
}

func (tok Token) IsScalar() bool {
	switch tok {
	case INT:
		return true
	}
	return false
}

func (tok Token) IsUnary() bool {
	switch tok {
	case ADD, SUB:
		return true
	}
	return false
}

func (tok Token) IsLeftAssoc() bool {
	// 是否是左结合
	switch tok {
	case INT, ADD, SUB, MUL, QUO, REM:
		return true
	}
	return false
}

func (tok Token) IsRightAssoc() bool {
	// 是否是右结合
	switch tok {
	case POW:
		return true
	}
	return false
}

func (tok Token) String() string {
	if 0 <= tok && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return tokens[0]
}
