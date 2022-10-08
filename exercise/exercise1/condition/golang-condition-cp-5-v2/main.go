package main

import "fmt"

func TicketPlayground(height, age int) int {
	if age < 5 {
		return -1
	} else if age > 12 {
		return 100000
	} else if (age == 12) || height > 160 {
		return 60000
	} else if (age >= 10 && age < 12) || height > 150 {
		return 40000
	} else if (age >= 8 && age < 10) || height > 135 {
		return 25000
	} else if (age >= 5 && age < 8) || height > 120 {
		return 15000
	}
	return 0
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(155, 6))
}
