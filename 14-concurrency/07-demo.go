/* concurrent safe state manipulation */

package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Counter :", counter)
}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
	wg.Done()
}
