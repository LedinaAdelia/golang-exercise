package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else if score < 70 || absent >= 5 {
		return "tidak lulus"
	} else {
		fmt.Print("Nilai Tidak Valid")
	}

	return ""
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 5))
	fmt.Println(GraduateStudent(90, 0))
	fmt.Println(GraduateStudent(70, 0))
	fmt.Println(GraduateStudent(70, 5))
	fmt.Println(GraduateStudent(0, 0))
	fmt.Println(GraduateStudent(0, 8))
}
