package main

func main(){
	println(add(10,20))
}


// func add(a,b interface{})interface{}{  // it gives error and also output is different if we dont use result variable in main
	
// 	aint,aIsInt :=a.(int)
// 	bint,bIsInt :=b.(int)
	
// 	if aIsInt && bIsInt{
// 		return aint+bint
// 	}

// 	afloat,aIsFloat :=a.(float64)
// 	bfloat,bIsFloat :=b.(float64)
	
// 	if aIsFloat && bIsFloat{
// 		return afloat+bfloat
// 	}

// 	astring,aIsString :=a.(string)
// 	bstring,bIsString :=b.(string)
	
// 	if aIsString && bIsString{
// 		return astring+bstring
// 	}

// 	return nil
// }


func add[T int | float64 | string](a, b T) T {
	return a + b
}