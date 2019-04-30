package main
import "fmt"
func main(){
	names := [5]string{"chaitu","bindhu","Nidhi","darpad","anuradha"}
	fmt.Println(names)
	var test [3]string
	test[0] = "age"
	test[1] = "sex"
	fmt.Println(test)
	//fmt.Println("test[3] is empty string", test[2] == nil )
	fmt.Println("test[3] is empty string", test[2] == "" )
}