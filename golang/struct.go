package main

import "fmt"

// struct is structure or a collection of fields
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pa = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v.X, v.Y) // we can access individual struct values via dot .

	fmt.Println(v1, pa, v2, v3)
}
