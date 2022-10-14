package main

import (
	"fmt"
)

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name                     string
	BaseSalary, WorkingMonth int
}

type Senior struct {
	Name                     string
	BaseSalary, WorkingMonth int
	PerformanceRate          float64
}

type Manager struct {
	Name                              string
	BaseSalary, WorkingMonth          int
	PerformanceRate, BonusManagerRate float64
}

func (j *Junior) GetBonus() float64 {
	sal := float64(j.BaseSalary)
	work := float64(j.WorkingMonth)
	works := work / 12
	count := 1 * sal * works
	fmt.Println("hehe: ", count)
	return count
}

func (s *Senior) GetBonus() float64 {
	sal := float64(s.BaseSalary)
	work := float64(s.WorkingMonth)
	works := work / 12
	perfR := s.PerformanceRate * sal
	count := (2 * sal * works) + perfR
	return count
}

func (m *Manager) GetBonus() float64 {
	sal := float64(m.BaseSalary)
	work := float64(m.WorkingMonth)
	works := work / 12
	perfR := m.PerformanceRate * sal
	mngB := m.BonusManagerRate * sal
	count := (2 * sal * works) + perfR + mngB
	return count
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	res := 0.0
	for _, v := range employees {
		total := EmployeeBonus(v)
		res += total
	}
	return res // TODO: replace this
}

func main() {
	employee := []Employee{
		&Junior{
			Name:         "Senior Engineer",
			BaseSalary:   100000,
			WorkingMonth: 12,
		},
		&Junior{
			Name:         "Senior Engineer 2",
			BaseSalary:   120000,
			WorkingMonth: 4,
		},
	}

	res2 := TotalEmployeeBonus(employee)
	fmt.Println(res2)
}
