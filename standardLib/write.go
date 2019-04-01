package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// write string to buffer variable
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "World!")

	num, err := b.WriteTo(os.Stdout)

	fmt.Println(num, err)  // 12 bytes
}
