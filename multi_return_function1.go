package main

import "fmt"

func swap(x, y, z string)(string,string,string) {
	return z, y, x
}

func main() {
	a, b, c := swap("welcome", "chaitu","to go lang")
	fmt.Println(a, b, c)
}
