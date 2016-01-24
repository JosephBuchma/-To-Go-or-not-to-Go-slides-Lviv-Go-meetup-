package main

import "fmt"

var hello = fmt.Sprintf("Hello")

func init() {
	fmt.Println(hello + " init")
}

func main() {
	fmt.Println("main")
}
