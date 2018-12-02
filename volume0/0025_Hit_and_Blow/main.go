package main

import "fmt"

func main() {
	var a1, a2, a3, a4, b1, b2, b3, b4 int
	for {
		i, _ := fmt.Scan(&a1, &a2, &a3, &a4)
		if i == 0 {
			break
		}
		fmt.Scan(&b1, &b2, &b3, &b4)
		a := [4]int{a1, a2, a3, a4}
		b := [4]int{b1, b2, b3, b4}
		var hit, blow int

		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if a[j] == b[k] {
					if j == k {
						hit++
					} else {
						blow++
					}
				}
			}
		}
		fmt.Println(hit, blow)
	}
}
