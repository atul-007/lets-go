package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://atul-ranjan.netlify.app/"

func main() {
	fmt.Println("Web requests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	output := string(databyte)
	fmt.Println(output)
	defer response.Body.Close()
}
