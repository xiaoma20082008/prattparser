//
// File: ast.go
// Project: prattparser
// File Created: 2024-08-01
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2024-08-01 11:58:18
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

import "fmt"

type (
	Expr interface{}

	ScalarExpr struct {
		Expr

		Val string
		Typ Token
	}

	UnaryExpr struct {
		Expr

		Op Token
		X  Expr
	}

	BinaryExpr struct {
		Expr

		L  Expr
		Op Token
		R  Expr
	}
)

func (s ScalarExpr) String() string { return fmt.Sprintf("%s", s.Val) }

func (s UnaryExpr) String() string { return fmt.Sprintf("%s%s", s.Op, s.X) }

func (s BinaryExpr) String() string { return fmt.Sprintf("(%s %s %s)", s.L, s.Op, s.R) }
