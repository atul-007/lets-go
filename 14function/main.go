package main

import "fmt"

func main() {
	fmt.Println("functions in go")
	first()
	value := adder(2, 4)
	fmt.Println("the sum is ", value)
	prores := proadder(1, 2, 3, 4, 5, 6)
	fmt.Println("result is", prores)

}
func first() {
	fmt.Println("first func created")
}
func adder(val1 int, val2 int) int /*return type of the function*/ {
	return val1 + val2
}
func proadder(value ...int) int {
	total := 0
	for _, value := range value {
		total += value
	}
	return total
}
