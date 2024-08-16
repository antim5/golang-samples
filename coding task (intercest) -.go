coding task (intercest) - 

func main() {
	set1 := []int{18, 1, 5, 30}
	set2 := []int{30, 32, 5, 4, 18}
	// [30, 5, 18], note: order in the result is irrelevant
	fmt.Printf("%v\n", intersect(set1, set2))
	// set1 = []int{1, 1, 0, 0}
	// set2 = []int{1, 0, 0, 0, 0}
	// // [1, 0, 0], note: try to intersect all suitable pairs, order in the result is irrelevant
	// fmt.Printf("%v\n", intersect(set1, set2))
}

func intersect(set1, set2 []int) []int {
  valuesMap := make(map[int]bool)
  var commonValues []int
  for v :=  range set1 {
    if ok := valuesMap[v]; !ok {
       valuesMap[v] = true
    }
  }
  for v :=  range set1 {
    if ok := valuesMap[v]; ok {
      commonValues = append(commonValues,v)
    }
  }
	return commonValues
}


