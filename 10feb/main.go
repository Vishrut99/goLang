package main

import (
	"fmt"
	"reflect"
	"strings"
)

// a:=5     // Can't declare here with shortened

func main() {
	main2()
	fmt.Println(ans())
	fmt.Println(ans1())
	fmt.Println(*ans2())

	a := 3
	b := 3.3

	fmt.Printf("a: %8T %[1]v \n", a)
	fmt.Printf("b: %8T %[1]v \n", b) // 1 based index

	const (
		s = "好abcd" // but we did not declare it like no shortend???
		// g:="hy"  // error we cant do shortend in const
		// v=s[2:]    // slice operations on string is not compile time
	)
	fmt.Printf("String :%v ,length :%d  \n", s, len(s)) // string is byte strem so it converts into that and take its length ,chinese charcter is off 3 length in byte in rune its 1
	// fmt.Printf("String :%s ,length :%d  \n", s, len(s)) //same output??what is difference in v and s
	fmt.Printf("s: |%010T|%2[1]v| \n", []rune(s))
	fmt.Printf("s: |%010T|%2[1]v| \n", []byte(s))

	// Strings are immutable if we try add operation like + then give new memmory location as once string length changes it give us new location or bytes array
	Str := `yey
	  yey  y
	  ey`
	fmt.Println(strings.Fields(Str))
	fmt.Println(len(strings.Fields(Str)))
	fmt.Println(&Str)
	Str += "aa" // String address will not change because its header...address of byte array in that header will change
	fmt.Println(&Str)

	func() {
		fmt.Println(a)
		a := 10
		fmt.Println(a)
	}()

	fmt.Println(a)

	x := [...]int{1, 2, 3}
	y := [3]int{}

	println(reflect.TypeOf(x) == reflect.TypeOf(y)) // What is reflect? why we use that?

	fmt.Println(x)

	type z int
	var x1 z
	y1 := 12

	// x1=y1  //Error.......underlying structure are same but type name is different..though internally it points to int

	x1 = z(y1) // not error because here type convert from int (y1) to (z)x1 because here it checks underlying structure
	fmt.Println(x1)
}

func ans() int {
	x := 5

	defer func() {
		x++
	}()
	return x // it makes temporary copy of return so defer changes orginal x but temp copy is not changed which will return
}

func ans1() (x int) {
	x = 5

	defer func() {
		x++
	}()
	return x // named parameter... it initialized x with zero value then assign 5 then return execite and this time it send x
	// but before return value defer execute and it change x to ++ so return then return x=6...no temp copy variable

	// return   // we can write only return also because it initialize x on its own and return means return that x value
}

func ans2() *int {
	x := 5

	defer func() {
		x++
	}()
	return &x // we return pointer and though it is unnamed parameter it copy the address of x
	// so copied address also point at x so when defer changes x it appears
}
