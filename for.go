//For is the only loop in go Lang
package main
import "fmt"

func main() {
		age := map[string]int{}
	age["chaitu"] = 35
	age["anuradha"] = 61
	age["Bindhu"] = 33
	age["nidhi"] = 8
	age["darpad"] = 5

	for name, ages := range age {
		switch ages {
		case 1,2,3,4,5:
			fmt.Println(name, "age is below 5 hence no ticket is required")
		case 6,7,8,9,10,11,12:
			fmt.Println(name, "age is below 12 hence half ticket is required")
		case 33:
			fmt.Println(name, "age is teenage")
		case 35:
			fmt.Println(name, "Too youg to retire")
		default:
			fmt.Println(name, "sit back and relax")


		}
	}

}