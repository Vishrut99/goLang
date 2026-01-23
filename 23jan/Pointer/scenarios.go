package main

import "fmt"

func Scenario() {
    // Scenario A: Using loop variable directly
    var pointers1 []*int
    
    for i := 0; i < 5; i++ {
        pointers1 = append(pointers1, &i)
    }
    
    fmt.Println("Scenario A:")
    for _, ptr := range pointers1 {
        fmt.Print(*ptr, " ")
    }
    fmt.Println()
    
    
    // Scenario B: Creating new variable each iteration
    var pointers2 []*int
    
    for i := 0; i < 5; i++ {
        temp := i
        pointers2 = append(pointers2, &temp)
    }
    
    fmt.Println("Scenario B:")
    for _, ptr := range pointers2 {
        fmt.Print(*ptr, " ")
    }
    fmt.Println()
    
    
    // Scenario C: Checking memory addresses
    fmt.Println("\nScenario C - Memory addresses:")
    fmt.Println("Pointers1 addresses:")
    for idx, ptr := range pointers1 {
        fmt.Printf("Index %d: Address=%p, Value=%d\n", idx, ptr, *ptr)
    }
    
    fmt.Println("\nPointers2 addresses:")
    for idx, ptr := range pointers2 {
        fmt.Printf("Index %d: Address=%p, Value=%d\n", idx, ptr, *ptr)
    }
    
    
    // Scenario D: What about this modification?
    fmt.Println("\nScenario D - After modification:")
    
    var pointers3 []*int
    numbers := []int{10, 20, 30, 40, 50}
    
    for i := 0; i < len(numbers); i++ {
        pointers3 = append(pointers3, &numbers[i]) 
    }
    
    fmt.Println("Before changing numbers:")
    for _, ptr := range pointers3 {
        fmt.Print(*ptr, " ")
    }
    
    numbers[2] = 999
    
    fmt.Println("\nAfter changing numbers[2] to 999:")
    for _, ptr := range pointers3 {
        fmt.Print(*ptr, " ")
    }
    fmt.Println()
}