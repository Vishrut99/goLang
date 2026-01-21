package main

import "fmt"

func main() {

	fmt.Println("=== Example 1: Default switch (auto break) ===")
	x := 1
	switch x {
	case 1:
		fmt.Println("case 1 executed")
	case 2:
		fmt.Println("case 2 executed")
	}
	// Output: only "case 1 executed"

	fmt.Println("\n=== Example 2: fallthrough ===")
	y := 1
	switch y {
	case 1:
		fmt.Println("case 1 executed")
		fallthrough
	case 2:
		fmt.Println("case 2 executed")
	case 3:
		fmt.Println("case 3 executed")
	}
	// Output:
	// case 1 executed
	// case 2 executed

	fmt.Println("\n=== Example 3: continue inside switch (continues LOOP) ===")

	for i := 1; i <= 3; i++ {
		fmt.Println("Loop iteration:", i)

		switch i {
		case 1:
			fmt.Println("  in case 1, continue loop")
			continue // continues FOR loop, not switch

		case 2:
			fmt.Println("  in case 2, continue loop")
			continue

		case 3:
			fmt.Println("  in case 3, end program")
		}

		fmt.Println("  this line runs only if no continue")
	}

	fmt.Println("\n=== Example 4: labeled continue from inside switch ===")

Outer:
	for i := 1; i <= 3; i++ {
		fmt.Println("Outer loop iteration:", i)

		for j := 1; j <= 3; j++ {
			fmt.Println("  Inner loop iteration:", j)

			switch j {
			case 2:
				fmt.Println("  j == 2, continue Outer loop")
				continue Outer // jumps to next iteration of OUTER loop
			}
		}

		fmt.Println("  this line runs only if inner loop not skipped")
	}
}
