package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

// buffer.go
type Buffer struct {
	io.Writer
}
func (b *Buffer) Write(in []bytes) { ... }

/*
Identify the RC for the following error:
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x46597d]
*/
func main() {
	var buf *bytes.Buffer
	write1(buf)
	if buf != nil {
		fmt.Println(buf)
	}
}

func write1(p io.Writer) {
	if p != nil {
		_, err := p.Write([]byte("test"))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
