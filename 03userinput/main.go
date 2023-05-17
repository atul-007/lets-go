package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := "welcome"
	fmt.Println(str)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza:")

	//comma ok //err err

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating: ", input)
	fmt.Printf("input is of type: %T ", input)

}
