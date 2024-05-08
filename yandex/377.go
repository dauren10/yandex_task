package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	var arr = make([][]int, a)
	un := make(map[int]int)
	for i := 0; i < a; i++ {
		temp := make([]int, 2)
		_, err := fmt.Scan(&temp[0], &temp[1])
		if err != nil {
			log.Fatal(err)
		}
		arr[i] = temp
	}

	fmt.Scan(&b)
	for i := 0; i < len(arr); i++ {
		_, v := arr[i][0], arr[i][1]
		un[v] += 1

	}
	result := []int{}
	for _, v := range un {
		result = append(result, v)
	}
	sort.Ints(result)
	if len(result) == a {
		fmt.Println(len(result))
	} else {
		fmt.Println(result[0])
	}

}
