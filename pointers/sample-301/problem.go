package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	var buf *bytes.Buffer
	boo(buf)
	if buf != nil {
		fmt.Println(buf)
	}
}

func boo(p io.Writer) {
	if p != nil {
		_, err := p.Write([]byte("test"))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
