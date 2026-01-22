package main

import "fmt"

func change(a [3]int) {
	a[0] = 999 // value change will not reflect outside
}

func main() {
	// arr := [3]int{1, 2, 3}
	// change(arr)
	// fmt.Println(arr)

	// a := make([]int, 0, 2)
	// b := a

	// // a = append(a, 1)
	// a = append(a, 2)

	// fmt.Println("Before:", a, b)

	// fmt.Println("Cap Before:", cap(a), cap(b))

	// a = append(a, 3) // may reallocate

	// b = append(b, 100) //	b overwrites a's value if no reallocation happened for a
	// b = append(b, 200)	// at this point if reallocation happened for a then b will have different underlying array
	// // b = append(b, 300)

	// fmt.Println("---------------------------------\nAfter:")
	// fmt.Println("a:", a) // a and b has same values because they point to same underlying array and b overwrites the values of a
	// fmt.Println("cap of a:", cap(a))
	// fmt.Println("pointer of a:", &a[0])// if no reallocation happened both a and b will have same pointer
	// fmt.Println("pointer of b:", &b[0])
	// fmt.Println("b:", b)
	// fmt.Println("cap of b:", cap(b))

	scores := [5]int{10, 20, 30, 40, 50}

	for i, v := range scores {
		scores[i] = v * 2 // changing value in array using index
		v = v * 3         // but changing v does not change the value in array because v is just a copy of the value in array
	}

	fmt.Println("Array after loop:", scores) // values in array are changed because we used index to change the value

	prices := []int{10, 20, 30, 40, 50}

	for i, v := range prices {
		prices[i] = v * 2
		v = v * 3
	}

	fmt.Println("Slice after loop:", prices)


	// Demonstrating append and its effect on underlying array

	count := 0

	for i,v := range prices {
		count = count + 1
		prices	[0] = 999
		prices = append(prices, i)
		fmt.Println(v)  //it will print the original value of slice before appending new value
		fmt.Printf("Iteration %d: prices = %v, count = %d\n", i, prices, count) // prices will show the updated slice after appending new value
	}

	fmt.Println("Final count:", count)
}
