package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("My modules")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("hey there mod users ")
}

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>i am god </h1>"))
}
