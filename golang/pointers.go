package main

import "fmt"

func main() {
	i, j := 10, 20
	fmt.Println(i)
	p := &i // point to i
	fmt.Println(*p)
	*p = 21        // set i through the pointer
	fmt.Println(i) // see new value of i
	fmt.Println(j)
}
