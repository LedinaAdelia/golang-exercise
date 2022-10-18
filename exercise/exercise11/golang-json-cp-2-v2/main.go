package main

import (
	"fmt"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
}



func LoanReport(data LoanData) LoanRecord {
	return LoanRecord{} // TODO: replace this
}

func RecordJSON(record LoanRecord, path string) error {
	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{"01-January-2021", []string{"1", "2"}},
			{"02-January-2021", []string{"1", "2", "3"}},
			{"03-January-2021", []string{"2", "3"}},
			{"04-January-2021", []string{"1", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
