package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}

	channel := make(chan int)

	res1 := do(channel)
	res2 := do1(res1)
	res3 := do2(res2)
	go func() {
		for _, v := range nums {
			channel <- v
		}
		close(channel)
	}()

	for res := range res3 {
		fmt.Println(res)
	}
}

func do(channel chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range channel {
			out <- v * 93
		}
		close(out)
	}()

	return out
}

func do1(channel chan int) (out chan int) {
	out = make(chan int)
	go func() {
		for v := range channel {
			out <- v + 107
		}
		close(out)
	}()
	return
}

func do2(channel chan int) (out chan int) {
	out = make(chan int)
	go func() {
		for v := range channel {
			out <- v & 10
		}
		close(out)
	}()
	return
}
