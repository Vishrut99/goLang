package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome!!!")

	text := "Hello, ?@#$%^&*()"
	for i, char := range text {
		fmt.Printf("Index: %d, Char: %c\n", i, char)
	}

	// a, f, c, d := 10, 20, 30, 40

	// {
	// 	fmt.Println(a, c, d, f)
	// 	b, c, d, e := "b", "c", 300, "400"
	// 	fmt.Println(a, b, c, d, f, e)

	// }

	// fmt.Println(a, c, d, f)
}
