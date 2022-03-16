package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan string, 3)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		ii := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- fmt.Sprintf("channel %d", ii)
		}()
	}

	go func() {
		for {
			select {
			case q := <-c:
				fmt.Println(q)
			}
		}
	}()

	wg.Wait()
	close(c)
}
