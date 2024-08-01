//
// File: scanner_test.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:44:21
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
	"reflect"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	type fields struct {
		src *bytes.Buffer
		nch rune
	}
	tests := []struct {
		name   string
		fields fields
		want   Lexed
	}{
		{"INT", fields{src: bytes.NewBuffer([]byte(`0 10 129 +1 -1`)), nch: 0}, Lexed{INT, "0"}},
		{"OPT", fields{src: bytes.NewBuffer([]byte(`+-*/%^()`)), nch: 0}, Lexed{ADD, "+"}},
		{"ILL", fields{src: bytes.NewBuffer([]byte(`a`)), nch: 0}, Lexed{ILLEGAL, "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				src: tt.fields.src,
				nch: tt.fields.nch,
			}
			if got := s.Scan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
