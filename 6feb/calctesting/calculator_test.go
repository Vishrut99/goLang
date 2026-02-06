package calculator

import "testing"


// ===== Add =====

func TestAdd_Parallel(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 6},
		{"zero", 0, 5, 5},
		{"negative", -2, -3, -5},
	}

	for _, tt := range tests {
		tt := tt // ⚠️ VERY IMPORTANT

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // makes this subtest parallel

			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}



// ===== Subtract =====
func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Error("Subtract failed")
	}
}


// ===== Multiply =====
func TestMultiply(t *testing.T) {
	if Multiply(4, 3) != 12 {
		t.Error("Multiply failed")
	}
}


// ===== Divide =====
func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)

	if err != nil {
		t.Error("Unexpected error")
	}

	if result != 5 {
		t.Error("Divide wrong result")
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}
}


// ===== IsPrime =====
func TestIsPrime(t *testing.T) {
	if !IsPrime(7) {
		t.Error("7 should be prime")
	}

	if IsPrime(9) {
		t.Error("9 should not be prime")
	}
}


// ===== Factorial =====
func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Error("Factorial failed")
	}

	if Factorial(-1) != 0 {
		t.Error("Negative factorial should be 0")
	}
}


// ===== Power =====
func TestPower(t *testing.T) {
	if Power(2, 3) != 8 {
		t.Error("Power failed")
	}
}


// ===== StringRepeat =====
func TestStringRepeat(t *testing.T) {
	if StringRepeat("a", 3) != "aaa" {
		t.Error("StringRepeat failed")
	}
}


// ===== Sum =====
func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	if Sum(nums) != 10 {
		t.Error("Sum failed")
	}
}


// ===== Average =====
func TestAverage_Table(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{"normal", []int{2, 4, 6}, 4},
		{"empty", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Average(tt.nums)
			if got != tt.want {
				t.Errorf("got %f want %f", got, tt.want)
			}
		})
	}
}

func TestAverageEmpty(t *testing.T) {
	if Average([]int{}) != 0 {
		t.Error("Average empty slice should be 0")
	}
}
