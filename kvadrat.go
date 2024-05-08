package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-3, 2, 4}
	result := make([]float64, len(arr))

	// Using two pointers to iterate through the array
	for i, j := 0, 0; i < len(arr); i, j = i+1, j+1 {
		result[j] = arr[i] * arr[i]
		fmt.Printf("%.0f ", result[j])
		math.Pow(arr[i], 2)
	}
}
