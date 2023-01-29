package main

import (
	"fmt"
)

func main() {
	mapA := map[int]string{101: "left", 102: "right"}
	mapB := map[int]string{201: "up", 202: "down"}

	fmt.Println(mergeMapsV2(mapA, mapB))
}

func mergeMapsV2[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
