package main

import "fmt"

func sum(x, y, z int) int {
	return x + y + z
}

func swap(x, y string) (string, string) {
	return y, x
}

func waris(total int) (son, daughter int) {
	son = total * 2 / 3
	daughter = total * 1 / 3
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(waris(9))
}
