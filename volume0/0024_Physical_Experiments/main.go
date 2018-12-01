package main

import "fmt"

func main() {
	var v float64
	for {
		i, _ := fmt.Scan(&v)
		if i == 0 {
			break
		}
		t := v / 9.8
		y := 4.9 * t * t

		N := 0
		for {
			N++
			if height := 5*N - 5; float64(height) > y {
				fmt.Println(N)
				break
			}
		}
	}
}
