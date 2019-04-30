package main
import "fmt"
func main() {
	ages := map[string]int{}
	ages["chaitu"] = 35
	switch {
	case ages["chaitu"] < 18:
		fmt.Println("chaitu can't vote")
	case ages["chaitu"] < 55:
		fmt.Println("chaitu can't retire")
	default:
		fmt.Println("he is enjoying")
	}
}