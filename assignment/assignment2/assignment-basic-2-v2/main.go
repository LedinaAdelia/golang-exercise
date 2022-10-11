package main

import (
	"fmt"
)

func GetDayDifference(date string) int {
	return 0 // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	return nil // TODO: replace this
}

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi helper untuk mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	return "" // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}
