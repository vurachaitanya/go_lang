package main
import "fmt"

func main() {
	fmt.Println(`
	Multi line string representation
	test
	`)
	fmt.Println("single line string")
	fmt.Println("\u2272")
	fmt.Println('a')
}

func mainn() {
	fmt.Println("comments defined") // in the same line

	//sample comments for single line
	
/* 

	Multi line comments

*/
}