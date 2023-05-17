/*
package main

import "fmt"

	func main() {
		defer fmt.Print("hello") //defer keyword moves the statement to the last line which leads to the last execution
		defer fmt.Println("bla1")
		defer fmt.Println("bla2") //and its first in last out so the statement deffered first will be executed last
		fmt.Println("world")
		atul()

}

	func atul() {
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}
	}
*/
package main

import "fmt"

func main() {
	fmt.Println("Hello,World!")
}
