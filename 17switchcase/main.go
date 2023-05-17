package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("dice game ")
	rand.Seed(time.Now().UnixNano())
	diceno := rand.Intn(6) + 1
	fmt.Println(diceno)

	switch diceno {
	case 1:
		fmt.Println("Dice value is one you can open ")
	case 2:
		fmt.Println("dice value is 2 you moved two spots")
	case 3:
		fmt.Println("dice value is 3 you moved three spots")
	case 4:
		fmt.Println("dice value is 4 you moved four spots")
	case 5:
		fmt.Println("dice value is 5 you moved five spots")
	case 6:
		fmt.Println("dice value is 6 roll again")
	default:
		fmt.Println("what was that?")
	}

}
