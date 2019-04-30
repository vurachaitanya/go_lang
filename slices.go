package main
import "fmt"
func main() {
	names := []string{}
	names = append(names, "chaitu")
	names = append(names, "bindhu", "darpad", "Nidhi", "anuradha")
	fmt.Println(names)
	fmt.Println("is the name empty:",names[3] == "")
}