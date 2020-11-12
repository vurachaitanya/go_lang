package main

import (
     "fmt" 
     "strconv"
)

func main() {
       var i int = 43
       fmt.Printf("%v , %T \n", i, i)

       var j string
       j = strconv.Itoa(i)
       fmt.Printf("%v, %T \n", j, j)
}

