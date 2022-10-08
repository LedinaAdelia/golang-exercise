package main

import "fmt"

func CountingNumber(n int) float64 {
	sum := 1.0
	total := 0.0
	for sum <= float64(n) {
		fmt.Println(sum)
		total += sum
		fmt.Println(total)
		fmt.Println("")
		sum += 0.5
	}
	return total // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(5))
}
