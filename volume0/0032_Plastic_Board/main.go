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
	var rectangle, diamond int
	for sc.Scan() {
		l := strings.Split(sc.Text(), ",")
		a, _ := strconv.Atoi(l[0])
		b, _ := strconv.Atoi(l[1])
		c, _ := strconv.Atoi(l[2])

		if a*a+b*b == c*c {
			rectangle++
		}
		if a == b {
			diamond++
		}
	}
	fmt.Println(rectangle)
	fmt.Println(diamond)
}
