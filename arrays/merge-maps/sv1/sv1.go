package main

import "fmt"

func main() {
	mapA := map[int]string{101: "left", 102: "right"}
	mapB := map[int]string{201: "up", 202: "down"}

	fmt.Println(mergeMapsV1(mapA, mapB))
}

func mergeMapsV1(ma map[int]string, mb map[int]string) map[int]string {
	for k, v := range mb {
		ma[k] = v
	}
	return ma
}
