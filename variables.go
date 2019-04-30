package main
import "fmt"
func main() {
	z := "test" // not declared and defined
	var k int //declared 
	k = 10 //defined 
	var a int = 16 //declared, Defined and assigned
	var b, c = "yes", true //string and boolean
	var d, e, _ = 12, 15, 11  //_ represents variable which declared but not used any where
	fmt.Println("variable declared are printed", a * 2)
	fmt.Println("print var", b, c)
	fmt.Println("sum of two variables", d + e)
	fmt.Println(z, k)

}