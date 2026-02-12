package main

import "sync"

func main() {
	// Copying a mutex is a common mistake that can lead to unexpected behavior and bugs in concurrent programs. When you copy a mutex, you create a new instance of the mutex that is not shared with the original. This means that the two mutexes will not synchronize access to shared resources correctly, leading to potential
	var mu sync.Mutex
	mu1 := mu
	b(mu1)

}

//go vet...............nocopy

func b(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	// Do some work while holding the lock
}
