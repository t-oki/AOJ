package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		var x1, y1, x2, y2, x3, y3 float64
		fmt.Scanf("%f%f%f%f%f%f", &x1, &y1, &x2, &y2, &x3, &y3)

		a := 2 * (x2 - x1)
		b := 2 * (y2 - y1)
		c := 2 * (x3 - x1)
		d := 2 * (y3 - y1)
		e := x2*x2 - x1*x1 + y2*y2 - y1*y1
		f := x3*x3 - x1*x1 + y3*y3 - y1*y1

		// 直線1: ax + by - e = 0
		// 直線2: cx + dy - f = 0
		px := (d*e - b*f) / (a*d - b*c)
		py := (c*e - a*f) / (b*c - a*d)
		r := math.Sqrt((px-x1)*(px-x1) + (py-y1)*(py-y1))
		fmt.Printf("%.3f %.3f %.3f\n", px, py, r)
	}
}
