package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlocks")

	myChannel := make(chan int, 2)

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myChannel)
	// myChannel <- 5
	wg.Add(2)
	//Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, res := <-myChannel
		fmt.Println(res)
		fmt.Println(val)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)
	//Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myChannel <- 0
		close(myChannel)
		// myChannel <- 6
		wg.Done()
	}(myChannel, wg)
	wg.Wait()

}
