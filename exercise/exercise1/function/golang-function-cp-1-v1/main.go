package main

import (
	"fmt"
	"time"
)

func DateFormat(day, month, year int) string {
	m := time.Month(month)
	d := dayFormatter(day)
	return fmt.Sprintf("%s-%s-%d", d, m, year)
}
func dayFormatter(day int) string {
	if day < 10 {
		return fmt.Sprintf("0%d", day)
	} else {
		return fmt.Sprintf("%d", day)
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(9, 2, 2012))
	// 08-February-2012
}
