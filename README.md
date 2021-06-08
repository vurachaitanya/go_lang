# go_lang


### Advantages of Go:
-	Code runs faster
-	Garbage collections
-	Simpler objects
-	Concurrency is inbuilt which improves efficiency using **goroutines
-	**Channels** help to solve this problem by providing a pattern that makes data safe from
concurrent modification. Channels help to enforce the pattern that only one goroutine should modify the data at any time.
  - The first **goroutine** passes a data value through the channel to a second goroutine that’s already waiting. The exchange of the data between both goroutines is synchronized, and once the hand-off occurs, both goroutines know the exchange took place.
- Go uses **composition** (inheritance vs composition) 
  -  Rather than building a long inheritance structure—Client extends User extends Entity—Go developers build small types—Customer and Admin and embed them into larger ones.
- Go has a modern **garbage collector** that does the hard work for you.
  - C++ & C, you need to allocate a piece of memory before you can use it, and then free it when you’re done. If you fail to do either of these correctly, you’ll have program crashes or memory leaks.
  


### Software Translation:
-	Machine language: Cpu instructions represented in binary
-	Assembly language: CPU instructions with Mnemonics - easy to read
-	High level Language: Commonly used language (C, C++, Java, Python, Go etc) - easy to use.



### Go Language
- Basics of GO Lang has been given each with an example 
- Variable, Strings, Conditions 

#### Learning Tools:
- Go-lang downloads: https://golang.org/
- Creating workspaces ,writing code and testing: https://golang.org/doc/code#Workspaces
- Various docs on go-lang: https://blog.golang.org/godoc
- [GO-by-Examples](https://gobyexample.com/)



### Bullet points 
- **imports** - imports code and give access to identifiers such as functions, constants and interfaces.
- log and OS are inbuilt function.
- A Unit compiled code is called **package.
- `_ "github.com/goinaction/code/chapter2/sample/matchers` - Technique to initialization from package even if you are not directly using, it will not import if not used. But init function calls if different code file within that package has it. - ONE WORD, imports if required only.
- **INIT** code would be called if defined before the main function. Even main is starting of the program but INIT will take the preceding.


### Variables :

###### GOPATH & GOROOT Should be set to execute your program
- For Variables which can directly used will get error. So we should use **map** and assign it to your variable
- All variables are initialized to their Zero Value. 
- `make(map[string]Matcher)` - A map is a reference type that you’re required to make in Go
- Numeric type value = 0
- Strings type value = empty string
- Booleans type value = false
- Pointers type value = nil
- reference type values = nil Zero - Nil

```
package main

import (
"log"
"os"

_ "github.com/goinaction/code/chapter2/sample/matchers"
"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
// Change the device for logging to stdout.
log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
// Perform the search for the specified term.
search.Run("president")
}
```
- ":=" This operator is used to both declare and initialize variables at the same time.
- Use the keyword go to launch and schedule goroutines to run concurrently. Anonymous function as a goroutine,  function that’s declared without a name is call Anonymous. In below code each feed be processed independently in a concurrent fashion. 

  - Functions accept a type Matcher and the address of a value type Feed. feed is a pointer variable. Pointer variables are grate to share variables between functions. They also allows function to access and change the state of the variable that we declared with in the scope of different functions and possible different goroutine.
  

```
go func(matcher Matcher, feed *Feed) {
  Match(matcher, feed, searchTerm, results)
  waitGroup.Done()
 }(matcher, feed)
```
