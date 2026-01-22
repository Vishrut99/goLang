package main

import "fmt"

func sum(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func main() {
	a := []int{1, 2, 3}

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(a...))

    work()// if you want to run work function from optimize.go
    work1()// if you want to run work1 function from optimize.go
    work2()// if you want to run work2 function from optimize.go
}

