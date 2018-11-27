package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	rs := 1
	for i := 1; i <= n; i++ {
		rs *= i
	}
	fmt.Println(rs)
}
