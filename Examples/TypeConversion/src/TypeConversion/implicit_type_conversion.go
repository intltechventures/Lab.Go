// An example of Implicit Type Conversion
// From a Twitter post by Dave DHeny
// https://twitter.com/davecheney/status/734646224696016901
// See:
// https://play.golang.org/p/0HfZFFnB7M
package main

import "fmt"

type stack []uintptr

// an anon type can be converted to a named type implicitly
func callers() stack {
	return make([]uintptr, 20)
}

func main() {
	fmt.Println(callers()[0])
}
