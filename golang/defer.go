package main

import "fmt"

func main() {
	// Does not call the func until all the surrouding func returns
	// first hello will get printed and then world, defer will hold the func call
	// Defer push the call to a stack and work as LIFO (last in first out)
	defer fmt.Println("world")
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End Statement")
}
