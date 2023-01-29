package main

import "fmt"

func main() {
	mapA := map[int]string{101: "left", 102: "right"}
	mapB := map[int]string{201: "up", 202: "down"}

	fmt.Println(mergeMaps(mapA, mapB))
	// output: map[101:left 102:right 201:up 202:down]
}

func mergeMaps(ma map[int]string, mb map[int]string) map[int]string {
	return nil
}
