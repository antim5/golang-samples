package main

import (
	"fmt"
)

/* Explain the func output */
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Iteration #", i)
		}()
	}
}
