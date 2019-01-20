package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		sc.Scan()
		line := sc.Text()

		insertCheck(line)
	}
}

func insertCheck(line string) {
	var B, C []int

	for _, v := range strings.Split(line, " ") {
		ball, _ := strconv.Atoi(v)
		switch {
		case len(B) == 0:
			B = append(B, ball)
		case B[len(B)-1] < ball:
			B = append(B, ball)
		case B[len(B)-1] >= ball:
			switch {
			case len(C) == 0:
				C = append(C, ball)
			case C[len(C)-1] < ball:
				C = append(C, ball)
			case C[len(C)-1] >= ball:
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
