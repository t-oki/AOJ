package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var x, y float64
	angle := 90.0

	for sc.Scan() {
		data := strings.Split(sc.Text(), ",")
		if data[0] == "0" && data[1] == "0" {
			break
		}
		meters, _ := strconv.Atoi(data[0])
		m := float64(meters)
		changedAngle, _ := strconv.Atoi(data[1])
		a := float64(changedAngle)

		radian := angle * math.Pi / 180
		x += m * math.Cos(radian)
		y += m * math.Sin(radian)
		angle -= a
	}
	fmt.Println(int(x))
	fmt.Println(int(y))
}
