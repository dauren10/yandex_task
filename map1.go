package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string][]string)
	m["Almaty"] = append(m["Almaty"], "Erevan", "Moscow")

	fmt.Println(m)

	arr := []string{"asdsad", "sadsad"}
	arr2 := []string{"yyyy", "ppp"}
	result := append(arr, arr2...)
	fmt.Println(result)

	g := [][]int{{1, 2}, {3, 4}, {1, 3}}
	sort.Slice(g, func(i, j int) bool {
		return g[i][0] < g[j][0]
	})

	fmt.Println(g)
}
