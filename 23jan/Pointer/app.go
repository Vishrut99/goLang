package main

// import "fmt"

// func modifySliceAppend(s []int) {
// 	s = append(s, 100)
// 	fmt.Println("Inside function - len:", len(s), "cap:", cap(s))
// }

// func modifySlicePointerAppend(s *[]int) {
// 	*s = append(*s, 200) // Modifies the ORIGINAL header!
// }

func main() {

	
	Scenario()

	// slice5 := make([]int, 3, 10)
	// slice5[0], slice5[1], slice5[2] = 1, 2, 3

	// fmt.Println("Before - len:", len(slice5), "cap:", cap(slice5))
	// fmt.Println("Slice visible elements:", slice5)

	// modifySliceAppend(slice5)

	// fmt.Println("After - len:", len(slice5), "cap:", cap(slice5))
	// fmt.Println("Slice visible elements:", slice5)

	// // The 100 is actually THERE in the underlying array!
	// // Proof: manually increase len and check
	// slice5 = slice5[:4] // Expand the slice view
	// fmt.Println("After expanding view - len:", len(slice5))
	// fmt.Println("Now we can see:", slice5) // [1 2 3 100] !!!



	// fmt.Println("----- Using Pointer to Slice -----")
	// slice4 := []int{1, 2, 3}
	// fmt.Println("Before - len:", len(slice4), "cap:", cap(slice4))
	// fmt.Println(slice4) // [1 2 3]

	// // Trying to append via function
	// modifySlicePointerAppend(&slice4)
	// fmt.Println("After - len:", len(slice4), "cap:", cap(slice4))
	// fmt.Println(slice4) // [1 2 3 200] ✓ Now it works!

	// With pointer:
	// ```
	// main's slice4:
	// ┌─────────────┐
	// │ ptr         │
	// │ len         │  ← Function has ADDRESS of this struct
	// │ cap         │  ← Can modify it directly!
	// └─────────────┘

}
