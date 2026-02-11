package main

import "fmt"

func main() {
	f := fib   // it is not closure is is just func pointer fp which points to fib
	g := fib() // g is closure...it points to anonymous function that fib  returns and it also has pointer to the a & b
	// since it is not declare or initialize in anonymous but it is in outer func fib
	// so once fib return anonymous func it also return env which have pointer to a and b ,,which escapes to heap
	//    g-> { fp           -- points to anonymous func
	//		  env          -- points to a and b
	//		}

	for i := 0; i < 10; i++ {
		fmt.Println((f())()) // why prints 1?
		// it will always return 1 because it is not closure and it does not have env to store a and b
	}
	fmt.Println("-------------")
	for i := 0; i < 10; i++ {
		fmt.Println(g()) // it will return the next fib number because it is closure and it has env to store a and b
	}

	fmt.Println("---------------")
	var v func()
	// i := 0 // why now 5
	for i := 0; i < 5; i++ {
		v = func() {
			fmt.Println(i)
		}

	}
	v()
	fmt.Println("------------------")

	i := 0
	for i = 0; i < 4; i++ {
		// j := i
		defer func() {
			fmt.Printf("%d @ %p\n", i, &i) // all print 4 because i is shared
		}() // but why not now
	}

}

func fib() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, b+a
		return b
	}
}
