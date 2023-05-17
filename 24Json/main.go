package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`              //to hide the password
	Tags     []string `json:"tags,omitempty"` //omitempty is used to remove the nil\null value
}

func main() {
	fmt.Println("Learning Json")

	//Encodejson()
	Decodejson()

}
func Encodejson() {

	courses := []Course{
		{"Reactjs bootcamp", 229, "blablabla.com", "abc123", []string{"web-dev", "js"}},
		{"mern bootcamp", 133, "blablabla.com", "abc123", []string{"full stack", "js"}},
		{"angular bootcamp", 229, "blablabla.com", "abc123", nil},
	}

	finaljson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)

}
func Decodejson() {
	jsondataformweb := []byte(`
	{
		 "coursename": "Reactjs bootcamp",
		"Price": 229,
		"Website": "blablabla.com",
		"tags": ["web-dev","js"]
	}`)

	var courses Course

	checkvalid := json.Valid(jsondataformweb)

	if checkvalid {
		fmt.Println("jsonn data is valid")
		json.Unmarshal(jsondataformweb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON DATA NOT VALID")
	}

	//Some cases where you just want to add data to the key value

	var myonlinedata map[string]interface{}

	json.Unmarshal(jsondataformweb, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for k, v := range myonlinedata {
		fmt.Printf("key is %v  and value is %v and type is %T\n", k, v, v)
	}
}
