package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var xa, ya, ra, xb, yb, rb float64
		fmt.Scan(&xa, &ya, &ra, &xb, &yb, &rb)
		distance := math.Sqrt((xb-xa)*(xb-xa) + (yb-ya)*(yb-ya))
		if rb+distance < ra {
			fmt.Println(2)
		} else if ra+distance < rb {
			fmt.Println(-2)
		} else if rb+ra >= distance {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
