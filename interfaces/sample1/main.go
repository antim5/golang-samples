package main

import "fmt"

type IA interface {
	Foo() string
}

type IB interface {
	Boo() string
}

type A struct{}

func (*A) Foo() string {
	return "Foo()"
}

type B struct{}

func (*B) Boo() string {
	return "Boo()"
}

type IAB interface {
	IA
	IB
}

type AB struct{}

func (*AB) Foo() string {
	return "Foo()"
}

func (*AB) Boo() string {
	return "Boo()"
}

type AB2 struct {
	A
	B
}

type AB3 struct {
	IA
	B
}

func main() {
	var a IA = &A{}
	var b IB = &B{}

	fmt.Println(a.Foo())
	fmt.Println(b.Boo())

	var ab IAB = &AB{}

	fmt.Println(ab.Foo())
	fmt.Println(ab.Boo())

	var ab2 IAB = &AB2{}

	fmt.Println(ab2.Foo())
	fmt.Println(ab2.Boo())

	ab22 := &AB2{}

	fmt.Println(ab22.A.Foo())
	fmt.Println(ab22.B.Boo())

	var ab3 IAB = &AB3{IA: &A{}}

	fmt.Println(ab3.Foo())
	fmt.Println(ab3.Boo())
}
