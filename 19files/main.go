package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in go")
	content := "i am god"
	file, err := os.Create("./input.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readfile("./input.txt")

}
func readfile(filename string) {
	readbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("reading file\n", readbyte)

}
