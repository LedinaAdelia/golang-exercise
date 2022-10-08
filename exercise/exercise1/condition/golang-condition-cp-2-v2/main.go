package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	result := 0.0
	if gender == "laki-laki" {
		result = (float64(height) - 100) - ((float64(height) - 100) * 0.1)
	} else if gender == "perempuan" {
		result = (float64(height) - 100) - ((float64(height) - 100) * 0.15)
	} else {
		fmt.Println("data tidak valid")
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
