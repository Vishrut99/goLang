package main

import (
	"fmt"
	"strings"
)

// we will optimize the log function by passing slice directly....3 allocations
func log1(level string, parts ...string) {
	message := level + ": " + strings.Join(parts, " ")
	fmt.Println(message)
}
func work1() {
	for i := 0; i < 3; i++ {
		log1("INFO", []string{"User", "logged", "in"}...)
		log1("DEBUG", []string{"Processing", "request", "ID", "12345"}...)
		log1("ERROR", []string{"Failed", "to", "connect"}...)
	}
}


// Another optimization by passing slice directly without variadic 3 allocations

func log2(level string, parts []string) {
	message := level + ": "
	for _, part := range parts {
		message = message + part + " "
	}
	fmt.Println(message)
}
func work2() {
	for i := 0; i < 3; i++ {
		log2("INFO", []string{"User", "logged", "in"})
		log2("DEBUG", []string{"Processing", "request", "ID", "12345"})
		log2("ERROR", []string{"Failed", "to", "connect"})
	}
}


// Original log function for comparison 4 allocations

func log(level string, parts ...string) {
	message := level + ": "
	for _, part := range parts {
		message = message + part + " "
	}
	fmt.Println(message)
}
func work() {
	for i := 0; i < 3; i++ {
		log("INFO", "User", "logged", "in")
		log("DEBUG", "Processing", "request", "ID", "12345")
		log("ERROR", "Failed", "to", "connect")
	}
}
