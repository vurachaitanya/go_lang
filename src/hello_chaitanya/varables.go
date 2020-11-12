package main

import "fmt"

/* Multi line comments
learn to explain more about your script
makes easy to read by a new Dev
*/

// test to comment single line comments.
var z int = 22

func main() {
        //var i int = 42
        
        i := 42
        k := 32.0

        //var i int
        //i = 42

        var j float32 = 27

	fmt.Println("Hello Chaitanya Vura",i) //trailing comments

	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%T",z)
	fmt.Println(z)
	fmt.Printf("%v, %T", k, k)

}
