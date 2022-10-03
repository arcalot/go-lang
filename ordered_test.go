package lang_test

import (
	"fmt"

	"go.arcalot.io/lang"
)

func compare[T lang.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func ExampleOrdered() {
	a := 1
	b := 2
	switch {
	case compare(a, b) < 0:
		fmt.Println("A is smaller than B")
	case compare(a, b) > 0:
		fmt.Println("B is smaller than A")
	default:
		fmt.Println("A and B are equal")
	}

	// Output: A is smaller than B
}
