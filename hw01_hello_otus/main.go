package main

import "golang.org/x/example/hello/reverse"

const message = "Hello, OTUS!"

func main() {
	println(reverse.String(message))
}
