// asmcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codegen

// Check small copies are replaced with moves.

func movesmall4() {
	x := [...]byte{1, 2, 3, 4}
	// 386:-".*memmove"
	// amd64:-".*memmove"
	// arm:-".*memmove"
	// arm64:-".*memmove"
	copy(x[1:], x[:])
}

func movesmall7() {
	x := [...]byte{1, 2, 3, 4, 5, 6, 7}
	// 386:-".*memmove"
	// amd64:-".*memmove"
	// arm64:-".*memmove"
	copy(x[1:], x[:])
}

func movesmall16() {
	x := [...]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	// amd64:-".*memmove"
	copy(x[1:], x[:])
}

// Check that no branches are generated when the pointers are [not] equal.

var x [256]byte

func ptrEqual() {
	// amd64:-"JEQ",-"JNE"
	// ppc64le:-"BEQ",-"BNE"
	// s390x:-"BEQ",-"BNE"
	copy(x[:], x[:])
}

func ptrOneOffset() {
	// amd64:-"JEQ",-"JNE"
	// ppc64le:-"BEQ",-"BNE"
	// s390x:-"BEQ",-"BNE"
	copy(x[1:], x[:])
}

func ptrBothOffset() {
	// amd64:-"JEQ",-"JNE"
	// ppc64le:-"BEQ",-"BNE"
	// s390x:-"BEQ",-"BNE"
	copy(x[1:], x[2:])
}
