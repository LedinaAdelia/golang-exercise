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
	General         Junior
	PerformanceRate float64
}

type Manager struct {
	General          Senior
	BonusManagerRate float64
}

func EmployeeBonus(employee Employee) float64 {
	return 0.0 // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	fmt.Println(employees)
	return 0.0 // TODO: replace this
}

func main() {
	Bonus := []Employee{
		Junior{
			Name:         "Della",
			BaseSalary:   1000000,
			WorkingMonth: 12,
		},
		Senior{
			General.Name:         "Ledina",
			General.BaseSalary:   100000,
			General.WorkingMonth: 12,
			PerformanceRate:      0.5,
		},
		Manager{
			General.Name:            "Adelia",
			General.BaseSalary:      100000,
			General.WorkingMonth:    12,
			General.PerformanceRate: 0.5,
			BonusManagerRate:        0.5,
		}
	}
	fmt.Println(Bonus)
}
