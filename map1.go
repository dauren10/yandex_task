package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["Almaty"] = append(m["Almaty"], "Erevan", "Moscow")

	fmt.Println(m)

	arr := []string{"asdsad", "sadsad"}
	arr2 := []string{"yyyy", "ppp"}
	result := append(arr, arr2...)
	fmt.Println(result)
}
