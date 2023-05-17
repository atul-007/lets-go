package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Handle Get and post Request ")
	//Performwebrequest()
	Performpostrequest()
}

func Performwebrequest() {
	const myurl = "https://645e3aae12e0a87ac0eadd29.mockapi.io/api/users"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status code", response.StatusCode)
	fmt.Println("content length is", response.ContentLength)

	var responsestring strings.Builder
	content, _ := io.ReadAll(response.Body)
	bytecount, _ := responsestring.Write(content)

	fmt.Println("bytecount", bytecount)
	fmt.Println(responsestring.String())
	//fmt.Println(string(content))
	defer response.Body.Close()
}

func Performpostrequest() {
	const myurl = "https://645e3aae12e0a87ac0eadd29.mockapi.io/api/users"
	//fake json payload
	requestbody := strings.NewReader(`
	{
		
		"name":"GAWD"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	output, _ := io.ReadAll(response.Body)
	fmt.Println(string(output))
}
