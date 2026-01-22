package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Write(value float64, file string) {
	text := fmt.Sprint(value)
	os.WriteFile(file, []byte(text), 0644)
}

func Read(file string) (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 1000, errors.New("Failed to read File")
	}
	ballanceText := string(data)
	balance, err := strconv.ParseFloat(ballanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to convert parse data")
	}
	return balance, nil
}
