package sample_301

import (
	"fmt"
	"strings"
)

func combine(element ...string) string {
	return strings.Join(element, ", ")
}

func main() {
	fmt.Println(combine())
	fmt.Println(combine("Red", "Yellow", "Green"))
	fmt.Println(combine("First", "Second"))
}
