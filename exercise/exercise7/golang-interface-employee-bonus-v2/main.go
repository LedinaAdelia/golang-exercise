package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func EmployeeBonus(employee Employee) float64 {
	return 0.0 // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	return nil // TODO: replace this
}

func main() {
	fmt.Println
}
