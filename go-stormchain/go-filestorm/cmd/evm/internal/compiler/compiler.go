// Copyright 2019 The go-filestorm Authors
// This file is part of go-filestorm.
//
// go-filestorm is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-filestorm is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-filestorm. If not, see <http://www.gnu.org/licenses/>.

package compiler

import (
	"errors"
	"fmt"

	"github.com/filestorm/go-filestorm/core/asm"
)

func Compile(fn string, src []byte, debug bool) (string, error) {
	compiler := asm.NewCompiler(debug)
	compiler.Feed(asm.Lex(src, debug))

	bin, compileErrors := compiler.Compile()
	if len(compileErrors) > 0 {
		// report errors
		for _, err := range compileErrors {
			fmt.Printf("%s:%v\n", fn, err)
		}
		return "", errors.New("compiling failed")
	}
	return bin, nil
}
