package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("enter the weight of the watermellon between 1 to  100")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		fmt.Println(err)

	} else {
		if num%2 == 0 && num != 2 {
			fmt.Println("yes")

		} else {
			fmt.Println("no")
		}
	}
}
