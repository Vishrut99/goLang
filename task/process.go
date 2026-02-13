package main

import (
	"fmt"
	"sync"
)

var eventMap = make(map[int]int)
var mu sync.Mutex

func process(a event) {
	// defer wg.Done()
	// time.Sleep(time.Second)
	mu.Lock()
	eventMap[a.userid] = a.Value
	id, _ := eventMap[a.userid]
	mu.Unlock()
	fmt.Println(":::::Stored:", id)
}
