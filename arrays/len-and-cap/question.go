package main

import "fmt"

func main() {
	init := [5]string{"A", "B", "C", "D"}
	s1 := init[2:5]

	fmt.Println(s1, len(s1), cap(s1))
}
