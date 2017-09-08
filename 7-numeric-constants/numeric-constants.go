package main

import "fmt"

const (
	Big   = 1 << 100 // 100000... s.d. 100x
	Small = 1 << 1   // 10
)

func main() {
	var i int = Small
	var f float64 = Big

	fmt.Printf("i bertipe %T dengan nilai %d\n", i, i)
	fmt.Printf("f bertipe %T dengan nilai %.2f\n", f, f)
}
