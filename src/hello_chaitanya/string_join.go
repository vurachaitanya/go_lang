package main

import (
	"fmt"
)

func main() {
	s := "You are been Hacked"
	s2 := "Please Pay Ransom"
	fmt.Printf("%v, %T \n", s, s2)
	fmt.Printf("%v, %T \n", s[2], s[2])
	fmt.Println(s + "," + s2)
}
