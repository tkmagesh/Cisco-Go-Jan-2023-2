/* concurrent safe state manipulation */

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter :", counter.count)
}
