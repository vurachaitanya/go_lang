package main

import (
	"fmt"
)

func main() {
	a := 1 + 2i
	b := 2 + 5.2i
	fmt.Printf("%v, %T \n", a, b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
