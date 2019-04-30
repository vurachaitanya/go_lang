package main
import "fmt"

func plus(a,b int) int {
	return a+b
}

func plusofplus(a,b,c int) int {
	return a+b+c
}

func main() {
	add := plus(12,44)
	fmt.Println("function output", add)

	add1 := plusofplus(11,12,13)
	fmt.Println("second function output", add1)
}