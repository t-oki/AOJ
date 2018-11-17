package main

import (
	"fmt"
)

func main() {
	for {
		var a int
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		rs := 0

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 10; k++ {
					for l := 0; l < 10; l++ {
						if i+j+k+l == a {
							rs++
						}
					}
				}
			}
		}
		fmt.Println(rs)
	}
}
