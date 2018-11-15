package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	for i := 0; i < l; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		c, _ := strconv.Atoi(s[2])

		n := []int{a, b, c}
		sort.Sort(sort.IntSlice(n))

		if square(n[0])+square(n[1]) == square(n[2]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func square(x int) int {
	return x * x
}
