package main
import "fmt"
func main() {
	age := map[string]int{}
	age["chaitu"] = 35
	switch age["chaitu"] {
	case 1,3,5,7,9,11,13:
		fmt.Println("chaitu's age is odd number age")
	case 2,4,6,8,10,12:
		fmt.Println("chaitu's age is even number age")
	default:
		fmt.Println("chaitu's is enjoying")
	}
}