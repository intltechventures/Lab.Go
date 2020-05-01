// Demonstration of the Constant Generator iota
package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota 
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 0..6
}



