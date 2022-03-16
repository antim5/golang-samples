package main

import "fmt"

type Stack struct {
	nums  []int
	mnums []int

	//nums // -1 -10 -5 -20 0 5 -100
	//mnums // -1 -10 -20 -100
}

func (s *Stack) Push(i int) {
	s.nums = append(s.nums, i)
	if len(s.mnums) == 0 {
		s.mnums = append(s.mnums, i)
	} else if i <= s.mnums[len(s.mnums)-1] {
		s.mnums = append(s.mnums, i)
	}
}

func (s *Stack) GetMin() (int, error) {
	if len(s.nums) == 0 {
		return 0, fmt.Errorf("empty")
	}

	/*	res := 100000 // maxint const
		for _, elem := range s.nums {
			if elem < res {
				res = elem
			}
		}*/

	return s.mnums[len(s.mnums)-1], nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.nums) == 0 {
		return 0, fmt.Errorf("empty")
	}

	res := s.nums[len(s.nums)-1]

	if res == s.mnums[len(s.mnums)-1] {
		s.mnums = s.mnums[:len(s.mnums)-1]
	}

	s.nums = s.nums[:len(s.nums)-1]

	return res, nil
}

func main() {
	/*
		MinStack minStack = MinStack()
		minStack.push(-10)
		minStack.push(0)
		minStack.push(-15)
		minStack.getMin() // return -15
		minStack.pop()
		minStack.getMin() // return -10
	*/

	s1 := &Stack{
		//minNum: 100000, // int.MAX
	}

	/*	fmt.Println(s1.GetMin())
		fmt.Println(s1.Pop())*/

	s1.Push(-2)
	s1.Push(0)
	s1.Push(-3)

	s1.Push(-3)
	s1.Push(-30)

	fmt.Println(s1.GetMin())
	fmt.Println(s1.Pop())

	fmt.Println(s1.GetMin())
	fmt.Println(s1.Pop())

	fmt.Println(s1.GetMin())
	fmt.Println(s1.Pop())

	fmt.Println(s1.GetMin())
	fmt.Println(s1.Pop())

	fmt.Println(s1.GetMin())
	fmt.Println(s1.Pop())

	fmt.Println(s1)
}
