package main

import (
	"example.com/mathlib"
	"example.com/msglb"
)

// global scoped variables
var (
	a = 10
	b = 20
)

func main() {
	// local scoped variables
	x := 5
	y := 10	
	msglb.Greeting()
	mathlib.Add(x, y)
	mathlib.Add(a, b)
	msglb.Goodbye()
}