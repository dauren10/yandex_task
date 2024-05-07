package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var arr = make([]int, 3)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	sort.Ints(arr)

	fmt.Println(arr[1])

}

func input_text() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	arr := []int{}
	inputText := scanner.Text()
	words := strings.Split(inputText, " ")
	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}

	sort.Ints(arr)

	fmt.Println(arr[1])
}

func by_one() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	arr := []int{}
	arr = append(arr, a, b, c)
	sort.Ints(arr)
	fmt.Println(arr[1])
}
