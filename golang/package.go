package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("This is package testing", rand.Int31)
	fmt.Println(math.Pi) // func start with capital letter are exported func from a package
}
