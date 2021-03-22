package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "Hello")
}
