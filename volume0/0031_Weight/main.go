package main

import (
	"fmt"
	"math"
)

func solve(n int) int {
	i := 0
	for {
		if val := math.Pow(2, float64(i)); val*2 > float64(n) {
			return int(val)
		}
		i++
	}
}

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		res := []int{}
		for {
			val := solve(n)
			res = append(res, val)
			n -= val

			if n <= 0 {
				break
			}
		}

		for i := range res {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(res[len(res)-i-1])
		}
		fmt.Println()
	}
}
