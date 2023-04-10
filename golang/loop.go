package main

import "fmt"

// Condition or loop statement go require expression surrounded by braces {}
func con(i int) int {
	if i > 10 {
		return i
	} else {
		return 0
	}
}
func main() {
	sum := 0
	// For loop, for is while in go if used without semicolons ;
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(con(11))
}
