package main

import "fmt"

const Pi = 3.14 // Const can be defined at package level or at func level and the name should start with capital letter

const (
	Big   = 1 << 100 // create a long number followed by 1 upto 100 zeros or different
	Small = 1 >> 99
)

func needFloat(x float32) float32 {
	return x * 0.1
}

func needInt(x int) int { return x*10 + 1 }

func main() {
	const Car = "Tata"
	fmt.Println("Happy", Pi, "day")
	fmt.Println(Car)
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
}
