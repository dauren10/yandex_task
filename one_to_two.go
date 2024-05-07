package main

import (
	"fmt"
)

func main() {
	// Исходный массив из 9 чисел
	numbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Создаем массив из трех подмассивов
	var result [3][3]int

	// Заполняем подмассивы
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			result[i][j] = numbers[i*3+j] //здесь не сумма а позиция в масссиве numbers
			fmt.Printf("%d * 3 + %d", i, j)
			fmt.Println()
		}
	}

	// Выводим результат
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			fmt.Println(result[i][j])
		}
	}
}
