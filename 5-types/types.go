package main

import "fmt"

func main() {
	// Numerik
	i := 1
	f := 2.45
	c := 0.17 + 0.5i
	i = int(f)

	// Boolean
	var b bool

	// String
	var s string

	fmt.Printf("%v %v %v %v %q\n", i, f, c, b, s)
}
