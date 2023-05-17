package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	num, _ := read.ReadString('\n')

	numrating, _ := strconv.ParseInt(strings.TrimSpace(num), 0, 64)
	var i int64

	//arr := make([]string, numrating)
	for i = 0; i < numrating; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		//arr[i] = input
		//chars := []rune(input)

		length := len(input)
		if length > 12 {
			m := length - 4
			l := strconv.Itoa(m)
			s := string(input[0]) + l + string(input[length-3])
			//arr[i] = s
			fmt.Println(strings.TrimSpace(s))
		} else {
			fmt.Println(strings.TrimSpace(input))
		}
	}
	/*for i := 0; i < len(arr); i++ {
		fmt.Println(strings.TrimSpace(arr[i]))
	}*/

}
