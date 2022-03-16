package main

import "fmt"

func findSumPair(nums []int, k int) (r []int) {
	sm := make(map[int]int)

	for idx, num := range nums {
		if _, ok := sm[num]; !ok {
			sm[num] = idx
		}
	}

	for _, num := range nums {
		if _, ok := sm[(k - num)]; ok {
			return []int{num, k - num}
		}
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
	return nil
}

func main() {
	/*
		findSumPair([2, -4, 10], 6) => [-4, 10]
		findSumPair([3, 6, 3], 6) => [3, 3]
		findSumPair([3, 6], 6) => None
	*/
	fmt.Println(findSumPair([]int{2, -4, 10}, 6))
	fmt.Println(findSumPair([]int{3, 6, 3}, 6))
	fmt.Println(findSumPair([]int{3, 6}, 6))
}
