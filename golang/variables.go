package main

import "fmt"

var c, p, java bool
var k, j int = 1, 2

func main() {
	var i int
	var fruit = "apple" // If initializer is present then the type can be ommited. the var will auto take the take based on the initializer
	car := "Tata"       // Short assignment, we can use := and skip var keyword inside a func (Note:-  it will not work outside of a func)
	model := 2023
	var a int
	var b string
	var c bool
	var d float32
	var model2 float32 = float32(model)
	v := 42 // Change value
	fmt.Println(i, c, p, java)
	fmt.Println(k, j, fruit)
	fmt.Println(car, model, model2)
	fmt.Println(a, b, c, d)            // Undefined variable will take default values like 0, "", false, 0
	fmt.Printf("v is of type %T\n", v) // TO check the var type we use Printf with %T
}

// BASIC TYPES IN GO

// bool, string

// int - negative + positive values
// uint - only positive values including 0

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128
