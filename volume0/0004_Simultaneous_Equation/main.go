package main

import (
	"fmt"
)

func main() {
	for {
		var a, b, c, d, e, f float64
		i, _ := fmt.Scan(&a, &b, &c, &d, &e, &f)
		if i == 0 {
			break
		}

		y := (c*d - a*f) / (b*d - a*e)
		x := (c - b*y) / a

		fmt.Printf("%.3f %.3f\n", x, y)
	}
}
