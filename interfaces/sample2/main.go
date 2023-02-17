package main

import (
	"context"
	"fmt"
)

const contextVar = "obj"

type Base interface {
	Base() string
}
type base struct{}

func (*base) Base() string {
	return "Base method is run"
}

func putBaseObject(ctx context.Context, object Base) context.Context {
	return context.WithValue(ctx, contextVar, object)
}

func getBaseObject(ctx context.Context) Base {
	return ctx.Value(contextVar).(Base)
}

type Extended interface {
	Base
	Extended() string
}
type extended struct{}

func (*extended) Extended() string {
	return "Extended method is run"
}

func (e *extended) Base() string {
	return e.Base()
}

func putExtendedObject(ctx context.Context, object Extended) context.Context {
	return context.WithValue(ctx, contextVar, object)
}

func getExtendedObject(ctx context.Context) Extended {
	return ctx.Value(contextVar).(Extended)
}

func main() {
	ctx := context.Background()

	sample := &extended{}
	ctx = putExtendedObject(ctx, sample)

	fmt.Printf("sample if to get as extended: %#v\n", getExtendedObject(ctx))
	fmt.Printf("sample if to get as base: %#v\n", getBaseObject(ctx))

	updatedSample := getBaseObject(ctx)
	ctx = putBaseObject(ctx, updatedSample)

	fmt.Printf("updated sample if to get as extended: %#v\n", getExtendedObject(ctx))
	fmt.Printf("updated sample if to get as base: %#v\n", getBaseObject(ctx))

	checkSample := getExtendedObject(ctx)
	fmt.Println(checkSample.Extended())
}
