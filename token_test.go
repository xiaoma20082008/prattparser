//
// File: token_test.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:22:31
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

import "testing"

func TestToken_Precedence(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want int
	}{
		{"+", ADD, 4},
		{"-", SUB, 4},
		{"*", MUL, 5},
		{"/", QUO, 5},
		{"^", POW, 6},
		{"%", REM, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tok.Precedence(); got != tt.want {
				t.Errorf("Token.Precedence() = %v, want %v", got, tt.want)
			}
		})
	}
}
