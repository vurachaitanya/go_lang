package main

import "fmt"

func main() {
        ages := map[string]int{}
        ages["Kevin"] = 57

        if ages["Kevin"] < 18 {
                fmt.Println("Kevin can't vote yet")
        } else if ages["Kevin"] < 67 {
                fmt.Println("Kevin is not of retirement age just yet")
        } else {
                fmt.Println("Enjoy retirement Kevin!")
        }
}