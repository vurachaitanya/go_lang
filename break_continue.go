package main
import "fmt"

func main() {
	a := 0
	for a < 10 {
		if a % 2 == 0 {
			a++
			continue
		} else if a == 5 {
			break
		} 
		fmt.Println("we are counting:",a)
		a++
		}
	}