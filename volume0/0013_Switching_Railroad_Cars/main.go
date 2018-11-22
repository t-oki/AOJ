package main

import "fmt"

func main() {
	stack := []int{}
	for {
		var a int
		i, _ := fmt.Scan(&a)
		if i == 0 {
			break
		}
		if a == 0 {
			fmt.Println(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, a)
	}
}
