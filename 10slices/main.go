package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")
	var slic = []string{"agffkf", "ajfgjf", "hfgaf"}

	slic = append(slic, "jfgkfkhf", "kwhfkwhgt")
	fmt.Println(slic)
	//slic = append(slic[:3])
	fmt.Println(slic)
	score := make([]int, 4)
	score[0] = 123
	score[1] = 1234
	score[2] = 12345
	score[3] = 123456
	fmt.Println(score)
	score = append(score, 12, 33)
	fmt.Println(score)
	sort.Ints(score)
	fmt.Println("sorted", score)
	fmt.Println(sort.IntsAreSorted(score))

	//to delete an element
	var str = []string{"atul", "akul", "kapil", "ayush", "jasghdjahgf", "sfjhaskfh"}

	str = append(str[:2], str[4:]...)

	fmt.Println(str)
}
