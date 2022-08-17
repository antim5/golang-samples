package main

import (
	"context"
	"fmt"
)

// Question: what is the app output?
func main() {
	contextA := context.TODO()
	contextB := context.WithValue(contextA, "key1", "value1")
	contextD := context.WithValue(contextA, "key2", "value2")
	fmt.Println(contextB.Value("key1"))
	fmt.Println(contextD.Value("key1"))
	fmt.Println(contextB.Value("key2"))
	fmt.Println(contextD.Value("key2"))
}
