package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is experimental repo, have fun!")
	for i := 0; i < 10; i++ {
	    fmt.Printf("i: %d\n", i);
	}
}

func forTest(a, b int) int {
	r := 10
	r = a + b*r
	if r == 10 {
		return r
	}
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	return r
}
