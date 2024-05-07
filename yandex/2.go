package main

import "fmt"

func main() {
	// Define the dimensions of the 2D array
	rows := 5
	cols := 5

	// Create the 2D array
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	// Print the array to verify
	for _, row := range arr {
		fmt.Println(row)
	}
}
