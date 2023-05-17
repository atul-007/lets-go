package main

import "fmt"

const Logintoken string = "jashbd" //capital first letter of the variable name to make it global it is equivalent to using the public keyword

func main() {
	var username string = "Atul Ranjan"
	fmt.Println(username)
	fmt.Printf("user name is of type : %T \n", username)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("isloggedin is of type : %T \n", isloggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("smallval is of type : %T \n", smallval)

	var val int = 255
	fmt.Println(val)
	fmt.Printf("val is of type : %T \n", val)

	var fracs float64 = 255.2464
	fmt.Println(fracs)
	fmt.Printf(" fracs is of type : %T \n", fracs)

	//implicit type

	var website = "atulranjan.com"
	fmt.Println(website)
	fmt.Printf(" website is of type : %T \n", website)

	//no var style

	bla := 545.5
	fmt.Println(bla)
	fmt.Printf(" bla is of type : %T \n", bla)

	fmt.Println(Logintoken)
	fmt.Printf(" Logintoken is of type : %T \n", Logintoken)

}
