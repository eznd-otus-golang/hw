package main

import "golang.org/x/example/hello/reverse"

const message = "Hello, OTUS!"

func main() {
	print(reverse.String(message))
}
