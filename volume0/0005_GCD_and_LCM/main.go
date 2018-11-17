package main

import (
	"fmt"
)

func main() {
	var a, b int
	for {
		i, _ := fmt.Scan(&a, &b)
		if i == 0 {
			break
		}

		gcd := getGCD(a, b)
		fmt.Println(gcd, getLCM(a, b, gcd))
	}
}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD(b, a%b)
}

func getLCM(a, b, gcd int) int {
	return a * b / gcd
}
