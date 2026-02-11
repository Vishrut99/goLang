package main

import "fmt"

func main() {
	// m := make(map[int]int)
	// m[1] = 1
	// m[2] = 2
	// m[3] = 3
	// fmt.Println("O.G.",m)
	// do(m)
	// fmt.Println("After do->", m)
	// dop(&m)
	// fmt.Println("After dop->", m)
	// fmt.Println(do())


	a:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("%v\n",a)
	b:=a[0:3]
	fmt.Printf("after %v\n",b)
	c:=b[0:5]
	fmt.Printf("after %v\n",c)
	d:=b[0:4:4]
	fmt.Println("d:",d,"cap of d=",cap(d))  // len=j-i cap=k-i ..d[i:j:k]
	d[0]=100
	fmt.Printf("after change first one a[%p]:%v\n",&a,a)
	fmt.Printf("after change first one d[%p]:%v\n",&d,d)
	d=append(d, 100)
	fmt.Printf("after append d: %v ..when we append in d it's underlying array change\n",d)
	fmt.Printf("after append a: %v .. it will not changed since d's underlying array changed\n",a)


	A:=[...]int{1,2,3,4,5,6,7,8,9}
	B:=A[0:4]
	D:=B[0:4:4]
	fmt.Printf("A: %v\n",A)
	fmt.Printf("D: %v\n",D)
	D[0]=100
	fmt.Printf("after change first one A:%v\n",A)
	fmt.Printf("after change first one D:%v\n",D)

	D=append(D, 100)
	fmt.Printf("after append D: %v ..when we append in D it's underlying array change\n",D)
	fmt.Printf("after append A: %v .. it will not changed since D's underlying array changed",A)

}

func do() (int){
	x:=1
	defer func() {
		x=2
	}()

	return x
}
// func do(m map[int]int) {
// 	m[4] = 4
// 	m = make(map[int]int)
// 	m[5] = 5
// 	fmt.Println("value ->", m)
// }

// func dop(m *map[int]int) {
// 	(*m)[6] = 6
// 	*m = make(map[int]int)
// 	(*m)[7] = 7
// 	fmt.Println("ptr ->", *m)
// }