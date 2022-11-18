package main

import (
	"fmt"
)

/* Explaint the func output */
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Iteration #", i)
		}()
	}
}
