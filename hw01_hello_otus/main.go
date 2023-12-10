package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

const message = "Hello, OTUS!"

func main() {
	fmt.Print(reverse.String(message))
}
