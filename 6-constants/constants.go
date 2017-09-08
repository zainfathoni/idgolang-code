package main

import "fmt"

const (
	Pi  = 3.14
	str = "Go \n\"keren\""
	c   = 'A'
	lit = `Go 
"keren"`
	b = true
)

func main() {
	fmt.Printf("%v %v %v\n", Pi, c, b)

	fmt.Println(str)
	fmt.Println(lit)
}
