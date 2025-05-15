package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //usually a pointer
var mut sync.Mutex    //usually a pointer

func main() {
	fmt.Println("Let's Learn GoRoutine")
	// go greeter("Hey")
	// greeter("Aruna")
	websitelist := []string{"https://example.com", "https://go.dev", "https://google.com", "https://github.com"}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond) //not an ideal way , we have packages for that like "sync"
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
	}
	fmt.Printf("%d  status code for website %s\n", res.StatusCode, endpoint)
}
