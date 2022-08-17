package main

import "fmt"

func main() {
	a := []int{10, 15, 35, 70, 105}
	reverse(a)
	fmt.Println(a)
	// Output: [105 70 35 15 10]
}
