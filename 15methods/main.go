package main

import "fmt"

func main() {

	atul := info{"Atul Ranjan", 19, "atulranjan@godev", "iit kgp"}
	fmt.Printf("details are %v", atul)
	atul.newuser()
	atul.newmail()
	fmt.Println(atul.email)

}

type info struct {
	Name    string
	Age     int
	email   string
	college string
}

func (i info) newuser() {
	fmt.Println("\n Name of the new user is ", i.Name)
}
func (f info) newmail() {
	f.email = "atulranjan@devgod"
	fmt.Println("New mail is ", f.email)
}
