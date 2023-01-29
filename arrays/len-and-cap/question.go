package main

import "fmt"

func main() {
	init := [5]string{"A", "B", "C", "D"}
	s1 := init[2:4]
	s2 := init[:3]
	s3 := init[3:]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
}
