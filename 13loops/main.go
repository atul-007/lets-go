package main

import "fmt"

func main() {
	fmt.Println("loops")
	week := []string{"monday", "tuesday", "wednesday", "Thrusday", "friday", "saturday", "sunday"}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	/*for i := range week {
		fmt.Println(i)

	}*/
	for index, day := range week {
		fmt.Printf("index in %v and day is %v", index, day)
		fmt.Println("")

	}
	rougevalue := 0
	for rougevalue < 10 {
		if rougevalue == 2 {
			goto ar
		}
		if rougevalue == 5 {
			break
		}
		if rougevalue == 3 {
			rougevalue++
			continue
		}
		rougevalue++
		fmt.Println(rougevalue)
	}
ar:
	fmt.Println("jumping to atulranjan.com")
}
