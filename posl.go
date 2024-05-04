package main

import "fmt"

func main() {
	source := "cartwheel"
	find := "car"

	res := fuzzysearch(find, source)
	fmt.Println(res)
}

// fuzzysearch('car', 'cartwheel');        // true
// fuzzysearch('cwhl', 'cartwheel');       // true
// fuzzysearch('cwheel', 'cartwheel');     // true
// fuzzysearch('cartwheel', 'cartwheel');  // true
// fuzzysearch('cwheeel', 'cartwheel');    // false
// fuzzysearch('lw', 'cartwheel');         // false
// найти из первого аргумента что есть во втором включая порядок

func fuzzysearch(find, source string) bool {
	for i := 0; i < len(source); i++ {
		fmt.Println(source[i])
	}
	return true
}
