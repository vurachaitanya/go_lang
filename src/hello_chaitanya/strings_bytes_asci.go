package main

import (
	"fmt"
)

func main() {
	s := "You are been Hacked"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
