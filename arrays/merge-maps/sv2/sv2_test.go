package main

import "testing"

var mapA = map[int]string{101: "left", 102: "right"}
var mapB = map[int]string{201: "up", 202: "down"}

func BenchmarkMergeMapsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeMapsV2(mapA, mapB)
	}
}
