package main

import "fmt"

type User struct {
	Name, Email string
}

func (u User) UserDetails() {
	fmt.Println("User name : ", u.Name)
	fmt.Println("User email : ", u.Email)
}

func main() {
	v := User{Name: "Arpit", Email: "arpit@gmail.com"}
	fmt.Println(v)
	v.UserDetails()
}
