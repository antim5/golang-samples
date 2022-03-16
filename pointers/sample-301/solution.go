package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	//var buf *bytes.Buffer - no init, pointer is nil
	buf := &bytes.Buffer{}

	//buf.WriteByte(0xff)

	foo(buf)

	if buf != nil {
		fmt.Println(buf)
	}
}

func foo(p io.Writer) {
	if p != nil {
		_, err := p.Write([]byte("test")) // nil pointer deref might happen here
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
