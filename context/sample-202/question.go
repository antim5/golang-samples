package main

import (
	"context"
	"fmt"
	"time"
)

// Question: what is the app output? for how long app runs?
func main() {
	sc := make(chan struct{})
	contextA, cancelContextA := context.WithCancel(context.Background())
	contextB, _ := context.WithTimeout(context.Background(), 4*time.Second)
  
	go func() {
		for {
			select {
				case <-contextA.Done():
					fmt.printline("A")
					sc <- struct{}{}
				case <-contextB.Done():
					fmt.printline("B")
					sc <- struct{}{}
			}
		}
	}()
  
	time.Sleep(2 * time.Second)
	cancelContextA()
	<-sc
}
