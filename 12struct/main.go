package main

import "fmt"

func main() {
	fmt.Println("struct")
	atul := User{"Atul Ranjan", 21212, "atulranjan789@gmail.com", 19, true}
	fmt.Println(atul)
	fmt.Printf("atul details are %+v\n", atul)
	fmt.Printf("My name is %v and email is %v", atul.name, atul.email)
}

type User struct {
	name   string
	rollno int
	email  string
	age    int
	status bool
}
