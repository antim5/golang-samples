package main

import "fmt"

func findSumPair(nums []int, k int) (r []int) {
	sm := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := sm[(k - num)]; ok {
			return []int{num, k - num}
		}
		sm[num] = struct{}{}
	}

	/*	for outerIdx, i := range nums {
			if outerIdx == len(nums)-1 {
				break
			}
			for _, j := range nums[outerIdx+1:] {
				if i+j == k {
					return []int{i, j}
				}
			}
		}
	*/
	return []int{}
}

func main() {
	/*
		findSumPair([2, -5, 5, 10],  5) =>  [-5, 10]
		findSumPair([3, 6, 3, 0],    6) =>  [3, 3]
		findSumPair([4, -4, 3, 10],  7) =>  [4, 3]
		findSumPair([3, 6],         15) =>  []
	*/
	fmt.Println(findSumPair([]int{2, -5, 5, 10}, 5))
	fmt.Println(findSumPair([]int{3, 6, 3, 0}, 6))
	fmt.Println(findSumPair([]int{4, -4, 3, 10}, 7))
	fmt.Println(findSumPair([]int{3, 6}, 15))
}
