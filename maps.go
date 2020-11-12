package main
import "fmt"
func main(){

	dateofbirth := map[string]string{
		"chaitu" : "20/04/1922",
		"bindhu" : "24/11/1920",
		"Nidhi"  : "02/10/2011",
		"darpad" : "08/08/2012",
		"anuradha" : "15/05/1911",
	}
	fmt.Println(dateofbirth)
	age := map[string]int{}
	age["chaitu"] = 22
	age["sri"] = 21
	age["nithya"] = 11
	age["tannu"] = 18
	age["chaitu"] = 6 //Latest defined would be given 
        fmt.Println(age["Nidhi"])
	fmt.Println(age)
	fmt.Println(age["chaitu"])
}
