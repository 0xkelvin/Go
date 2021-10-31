package main

import "fmt"

func main() {
	vair i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)	
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
