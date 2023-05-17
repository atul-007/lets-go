package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ksafjhds"

func main() {

	fmt.Println("URLS in go")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawPath)

	para := result.Query()
	fmt.Printf("Datatype of para %T\n", para)
	fmt.Println(para["coursename"])
	for _, val := range para {
		fmt.Println(val)
	}
	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user-atul",
	}
	newurl := partsofurl.String()
	fmt.Println(newurl)

}
