package main

import (
	"runtime"
)

func main() {

	m := make(map[int][]byte)

	for i := 0; i < 1_000; i++ {
		m[i] = make([]byte, 1_000_000)
	}

	for i := 0; i < 1_000; i++ {
		delete(m, i)   // does not free the backing arrays of the slices
	}

	// x := 10
	// ptr := unsafe.Pointer(&x)

	// _ = ptr // GC does NOT see unsafe.Pointer as a root

	// var leak []int

	// func() {
	// 	data := make([]int, 10_000_000)
	// 	leak = data[:1] // tiny slice, huge backing array
	// }()

	// fn := func() func() int {
	// 	big := make([]byte, 10_000_000)
	// 	return func() int {
	// 		return len(big)
	// 	}
	// }()

	runtime.GC()

	// fmt.Println("length: ", len(leak))
	// fmt.Println("capacity: ", cap(leak))

	// fmt.Println("fn(): ", fn())

	// fmt.Println(x)

}
