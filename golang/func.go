package main

import "fmt"

// If we are return something from func then we have to define the data type of that
func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // Naked return values
}

func main() {
	fmt.Println(add(5, 5))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(7))
}
