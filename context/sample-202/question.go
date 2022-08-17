package main

import (
	"context"
	"fmt"
	"time"
)

// Question: what is the app output? for how long it takes?
func main() {
	sc := make(chan struct{})
	contextA, canelContextA := context.WithCancel(context.Background())
	contextB, _ := context.WithTimeout(contextA, 4*time.Second)
	go func() {
		select {
		case <-contextA.Done():
			fmt.Println("A")
			sc <- struct{}{}
		case <-contextB.Done():
			fmt.Println("B")
			sc <- struct{}{}
		}
	}()
	time.Sleep(2 * time.Second)
	canelContextA()
	<-sc
}
