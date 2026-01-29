package main

import (
	"fmt"
)

func main() {

	// s := make([]func(), 4)

	i := 0
	for i = 0; i < 4; i++ {
		// j := i
		defer func() {
			fmt.Printf("%d @ %p\n", i, &i) // all print 4 because i is shared
		}() // but why not now
	}

	// for i := 0; i < 4; i++ {
	// 	s[i]()
	// }

	f, g := fib(), fib()

	fmt.Println(f(), f(), f(), g(), g()) // f and g are different closures so they maintain separate states
	fmt.Println(f(), f(), g(), g(), g())

}
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
