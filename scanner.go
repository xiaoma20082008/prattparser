//
// File: scanner.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:11:17
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

import "bytes"

type Scanner struct {
	src *bytes.Buffer
	nch rune
}

type Lexed struct {
	Tok Token
	Lit string
}

func (s *Scanner) Read() rune {
	var ch rune
	var err error
	if s.nch != 0 {
		ch = s.nch
		s.nch = 0
	} else if ch, _, err = s.src.ReadRune(); err != nil {
		return 0
	}
	return ch
}

func (s *Scanner) Scan() Lexed {
	ch := s.Read()
	for ch == ' ' || ch == '\r' || ch == '\n' || ch == '\t' || ch == '\f' {
		ch = s.Read()
	}
	var lexed Lexed
	switch true {
	case '0' <= ch && ch <= '9':
		buf := []rune{ch}
		ch = s.Read()
		for '0' <= ch && ch <= '9' {
			buf = append(buf, ch)
			ch = s.Read()
		}
		s.nch = ch
		lexed.Tok = INT
		lexed.Lit = string(buf)
	default:
		switch ch {
		case 0:
			lexed.Tok = EOF
			lexed.Lit = EOF.String()
		case '+':
			lexed.Tok = ADD
			lexed.Lit = ADD.String()
		case '-':
			lexed.Tok = SUB
			lexed.Lit = SUB.String()
		case '*':
			lexed.Tok = MUL
			lexed.Lit = MUL.String()
		case '/':
			lexed.Tok = QUO
			lexed.Lit = QUO.String()
		case '^':
			lexed.Tok = POW
			lexed.Lit = POW.String()
		case '%':
			lexed.Tok = REM
			lexed.Lit = REM.String()
		}
	}
	return lexed
}
