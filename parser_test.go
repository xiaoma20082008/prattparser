// File: parser_test.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 15:25:05
// Last Modified By: xiaoma20082008 (mmccxx2519@gmail.com>)
// ------------------------------------------------------------------------
//
// Copyright (C) xiaoma20082008. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package prattparser

import (
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name   string
		fields string
		want   Expr
	}{
		{"number", "-1", UnaryExpr{Op: SUB, X: ScalarExpr{Val: "1", Typ: INT}}},
		{"calculator", "-1+2*3", BinaryExpr{L: UnaryExpr{X: ScalarExpr{Val: "1", Typ: INT}, Op: SUB}, Op: ADD, R: BinaryExpr{L: ScalarExpr{Val: "2", Typ: INT}, Op: MUL, R: ScalarExpr{Val: "3", Typ: INT}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.fields)
			if got := p.Parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
