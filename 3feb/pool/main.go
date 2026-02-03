package main

import (
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("🆕 Creating new buffer")
		return make([]byte, 10)
	},
}

func main() {
	// First Get - pool is empty, calls New()
	buf1 := bufferPool.Get().([]byte)
	fmt.Println("Got buf1")

	// Second Get - pool still empty, calls New() again
	buf2 := bufferPool.Get().([]byte)
	fmt.Println("Got buf2")

	// Put buf1 back
	bufferPool.Put(buf1)
	fmt.Println("Put buf1 back")

	// Third Get - pool has buf1, reuses it! No New() call
	buf3 := bufferPool.Get().([]byte)
	fmt.Println("Got buf3 (reused buf1)")

	buf1[0] = 1
	buf2[0] = 2
	buf3[1] = 1

	// Put both back
	bufferPool.Put(buf2)
	bufferPool.Put(buf1)
	// bufferPool.Put(buf3)

	bufferPool.Put([]byte("hello"))
	bufferPool.Put([]byte{1, 2, 3, 4, 5})
	fmt.Println(bufferPool.Get().([]byte))
	// fmt.Println(bufferPool.Get().([]byte))
	fmt.Println(bufferPool.Get().([]byte))
	fmt.Println(bufferPool.Get().([]byte))
	fmt.Println(bufferPool.Get().([]byte))
}
