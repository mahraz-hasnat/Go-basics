package main

// global scoped variables
var (
	a = 10
	b = 20
)

func main() {
	// local scoped variables
	x := 5
	y := 10	
	greeting()
	add(x, y)
	add(a, b)
	goodbye()
}