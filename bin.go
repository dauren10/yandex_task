package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := bin(arr, 8)
	fmt.Println("res", res)
}

func bin(arr []int, n int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		mid := (min + max) / 2

		val := arr[mid]
		if val == n {
			return mid
		}

		if n > val {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1 // возвращаем -1, если элемент не найден
}
