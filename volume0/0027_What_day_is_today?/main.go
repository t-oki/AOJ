package main

import "fmt"

func main() {
	days := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weekday := [7]string{"Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday"}

	for {
		var m, d int
		fmt.Scan(&m, &d)
		if m == 0 {
			break
		}

		for i := 0; i < m-1; i++ {
			d += days[i]
		}

		fmt.Println(weekday[d%7])
	}
}
