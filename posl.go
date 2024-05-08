package main

import "fmt"

func fuzzysearch(sub string, s string) bool {
	subLen := len(sub)
	sLen := len(s)

	subPtr := 0
	sPtr := 0

	for subPtr < subLen && sPtr < sLen {
		if sub[subPtr] == s[sPtr] {
			subPtr++
		}
		sPtr++
	}

	return subPtr == subLen

}

func main() {
	fmt.Println(fuzzysearch("car", "cartwheel"))       // true
	fmt.Println(fuzzysearch("cwhl", "cartwheel"))      // true
	fmt.Println(fuzzysearch("cwheel", "cartwheel"))    // true
	fmt.Println(fuzzysearch("cartwheel", "cartwheel")) // true
	fmt.Println(fuzzysearch("cwheeel", "cartwheel"))   // false
	fmt.Println(fuzzysearch("lw", "cartwheel"))        // false
}
