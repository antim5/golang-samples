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

	sc := make(chan int)
	go func() {
		for q := range c {
			fmt.Println(q)
		}
		sc <- 1
	}()

	wg.Wait()
	close(c)

	<-sc
}
