package main

import (
	"fmt"
)

func intersect(set1, set2 []int) []int {
	elemsRegister := make(map[int]int)
	var result []int

	for _, elem := range set1 {
		if _, ok := elemsRegister[elem]; !ok {
			elemsRegister[elem] = 1
		} else {
			elemsRegister[elem] += 1
		}
	}
	for _, elem := range set2 {
		if count, ok := elemsRegister[elem]; ok && count > 0 {
			elemsRegister[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	set1 := []int{18, 1, 5, 30}
	set2 := []int{30, 32, 5, 4, 18}
	// [30, 5, 18], note: order in the result is irrelevant
	fmt.Printf("%v\n", intersect(set1, set2))
	set1 = []int{1, 1, 0, 0}
	set2 = []int{1, 0, 0, 0, 0}
	// [1, 0, 0], note: try to intersect all suitable pairs, order in the result is irrelevant
	fmt.Printf("%v\n", intersect(set1, set2))
}
