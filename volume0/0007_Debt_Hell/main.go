package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	r := 100.0
	for i := 0; i < n; i++ {
		r = math.Ceil(1.05 * r)
	}

	fmt.Println(int(r * 1000))
}
