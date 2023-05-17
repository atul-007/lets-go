package main

import "fmt"

func main() {
	fmt.Println("pointers")
	//var ptr *int
	user := 230498
	var ptr = &user
	fmt.Println("value of ptr is ", *ptr)
	fmt.Println("address of ptr is ", ptr)
}
