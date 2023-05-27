package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}
	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("one R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.RLock()
		score = append(score, 3)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
}
