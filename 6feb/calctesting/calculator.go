package calculator

import (
	"errors"
	"fmt"
	// "fmt"
	"math"
)

func main() {
	// Example usage
	a, b := 10, 5
	fmt.Printf("Add: %d + %d = %d\n", a, b, Add(a, b))
	fmt.Printf("Subtract: %d - %d = %d\n", a, b, Subtract(a, b))
	fmt.Printf("Multiply: %d * %d = %d\n", a, b, Multiply(a, b))
	if result, err := Divide(float64(a), float64(b)); err == nil {
		fmt.Printf("Divide: %d / %d = %.2f\n", a, b, result)
	} else {
		fmt.Printf("Divide error: %s\n", err)
	}
	fmt.Printf("IsPrime(%d) = %t\n", a, IsPrime(a))
	fmt.Printf("Factorial(%d) = %d\n", a, Factorial(a))
	fmt.Printf("Power(%d, %d) = %d\n", a, b, Power(a, b))
	fmt.Printf("StringRepeat(\"Go\", %d) = %s\n", b, StringRepeat("Go", b))
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum(%v) = %d\n", numbers, Sum(numbers))
	fmt.Printf("Average(%v) = %.2f\n", numbers, Average(numbers))
}

// Add returns the sum of two numbers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two numbers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the division of two numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Factorial calculates the factorial of a number
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Power calculates base raised to the power of exponent
func Power(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

// StringRepeat repeats a string n times
func StringRepeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// Sum calculates the sum of a slice of integers
func Sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Average calculates the average of a slice of integers
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Sum(numbers)) / float64(len(numbers))
}
