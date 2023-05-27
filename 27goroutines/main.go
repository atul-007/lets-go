package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	fmt.Println("Go Routines")

	websitelist := []string{
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
		"https://gthub.com",
		"https://atul-ranjan.netlify.app/",
	}
	for _, web := range websitelist {
		wg.Add(1)
		getstatuscode(web)

	}
	wg.Wait()
	fmt.Println(signals)
}
func getstatuscode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
