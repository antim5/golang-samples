package main

import "fmt"

// implement the 'combine' function
func main() {
	fmt.Println(combine())
	// output: none

	fmt.Println(combine("Red", "Yellow", "Green"))
	// output: "Red Yellow Green"

	fmt.Println(combine("First", "Second"))
	// output: "First Second"
}

func combine(s ...string) string {
	return s[0]
}
