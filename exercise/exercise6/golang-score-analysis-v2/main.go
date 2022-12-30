package main

import (
	"fmt"
	"sort"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	if grades != nil {
		s.Grades = grades
	}
}

func Analysis(s School) (float64, int, int) {
	a := 0.0

	if len(s.Grades) == 0 {
		return 0.0, 0, 0
	} else {
		sort.Ints(s.Grades)
	}

	for _, v := range s.Grades {
		a += float64(v)
	}

	Average := a / float64((len(s.Grades)))
	Min := s.Grades[0]
	Max := s.Grades[len(s.Grades)-1]

	return Average, Min, Max

}

func main() {
	fmt.Println(Analysis(School{}))
}
