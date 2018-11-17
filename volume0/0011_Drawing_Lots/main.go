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
	sc.Scan()
	lines, _ := strconv.Atoi(sc.Text())
	l := make([]int, lines)
	for i := 0; i < lines; i++ {
		l[i] = i + 1
	}

	sc.Scan()
	rows, _ := strconv.Atoi(sc.Text())
	for i := 0; i < rows; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), ",")
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])

		l[a-1], l[b-1] = l[b-1], l[a-1]
	}
	for _, j := range l {
		fmt.Println(j)
	}
}
