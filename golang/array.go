package main

import "fmt"

func main() {
	// Array in go are of fixed size
	fmt.Println("Array")
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a, a[0], a[1])
	primes := [6]int{1, 2, 3, 4, 5, 6}
	b := [1]int{}
	b[0] = 1
	fmt.Println(b)
	fmt.Println(primes)
}
