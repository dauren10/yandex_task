package main

import "fmt"

func main() {
	mystr := "dauren"
	for i := 0; i < len(mystr); i++ {
		fmt.Println(string(mystr[i]))
	}
}
