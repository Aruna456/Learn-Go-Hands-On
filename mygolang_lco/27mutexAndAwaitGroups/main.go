package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Let's Learn about RACE CONDITION!")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	//immediately invoked funcs
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Function 1")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Function 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Function 3")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		// m.Lock()
		// score = append(score, 3)
		// m.Unlock()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
