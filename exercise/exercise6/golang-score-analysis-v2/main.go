package main

import (
	"fmt"
	"sort"
	"strconv"
)

// func Add(a, b int) int {
// 	return 0
// }

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	if len(grades) == 0 {

	} else {
		s.Grades = grades
	}
}

func Analysis(s School) (float64, int, int) {
	s.AddGrade(s.Grades...)
	fmt.Println(s.Grades)
	count := 0
	avg := 0.0
	min := 0
	max := 0
	if len(s.Grades) == 0 {

	} else {
		sort.Slice(s.Grades, func(i, j int) bool {
			return s.Grades[i] < s.Grades[j]
		})
		for _, v := range s.Grades {
			count += v
		}
		min = s.Grades[0]
		max = s.Grades[(len(s.Grades) - 1)]
		try := float64(count) / float64(len(s.Grades))
		a := fmt.Sprintf("%.2f", try)
		avg, _ = strconv.ParseFloat(a, 64)
	}
	if len(s.Grades) == 1 {
		s.Grades = append(s.Grades, min, max)
	}
	return avg, min, max

}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60, 60, 100, 100, 100, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57, 76, 87, 89, 54, 43, 12, 15, 16, 17, 100, 80, 87, 86, 57, 57},
	})

	fmt.Println(avg, min, max)
}
