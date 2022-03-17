package main

import (
	"context"
	"fmt"
)

func main() {
	contextA := context.Background()
	contextB := context.WithValue(contextA, "key1", "value1")
	contextC, cancelContextC := context.WithCancel(contextB)
	contextD := context.WithValue(contextC, "key2", "value2")

	fmt.Println(contextB.Value("key1"))
	fmt.Println(contextD.Value("key1"))

	fmt.Println(contextB.Value("key2"))
	fmt.Println(contextD.Value("key2"))
	
	cancelContextC()
	fmt.Println(contextD.Value("key2"))
}
