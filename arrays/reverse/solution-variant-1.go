package main

import "fmt"

func main() {
	a := []int{10, 15, 35, 70, 105}
	reverse(a)
	fmt.Println(a)
	// Output: [105 70 35 15 10]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
