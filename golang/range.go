package main

import "fmt"

var pow = []int{1, 2, 3, 4, 4, 5, 6, 67}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
