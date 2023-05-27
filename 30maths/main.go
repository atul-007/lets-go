package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
	//"time"
)

func main() {

	//rand.Seed(time.Now().UnixNano()) //generating random num through math package
	//fmt.Println(rand.Intn(5))

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5)) //generating  random num  through crypto
	fmt.Println(myRandomNum)
}
