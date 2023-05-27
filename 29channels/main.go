package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in go")

	mych := make(chan int, 2) //create a channel with boundary 2
	wg := &sync.WaitGroup{}

	wg.Add(2)

	//read only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, ischanelopen := <-mych //check if channel is open and pass its value to val
		fmt.Println(ischanelopen)
		fmt.Println(val)

		wg.Done()

	}(mych, wg)

	//send only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		mych <- 0   //add value to the channel
		close(mych) //close channel
		wg.Done()

	}(mych, wg)
	wg.Wait()
}
