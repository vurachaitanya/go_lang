package main

import (
	"fmt"
)

func main() {
	s := "You are been Hacked"
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", s[2], s[2])
	fmt.Printf("%v, %T \n", string(s[2]), s[2])
}
