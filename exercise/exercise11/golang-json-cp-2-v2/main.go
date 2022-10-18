package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
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
	MonthDate    string     `json:"month_date"`
	StartBalance float64    `json:"start_balance"`
	EndBalance   float64    `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	Id         string  `json:"id"`
	Total_loan float64 `json:"total_loan"`
}

func LoanReport(data LoanData) LoanRecord {
	recLoan := LoanRecord{}
	var loan float64 = 50000
	try := map[string]float64{}
	sBal := data.StartBalance
	date := ""
	for _, a := range data.Data {
		date = a.Date
		for _, b := range a.EmployeeIDs {
			if sBal <= 0 {
				break
			} else if sBal > 1 {
				try[b] += loan
				sBal -= int(loan)
			}
		}
	}

	keys := make([]string, 0, len(try))
	for k := range try {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fixDate := ""
	if date == "" {

	} else {
		formatDate := strings.Split(date, "-")
		fixDate = formatDate[1] + " " + formatDate[2]
	}
	recLoan.MonthDate = fixDate
	recLoan.StartBalance = float64(data.StartBalance)
	recLoan.EndBalance = float64(sBal)

	for _, v := range keys {
		recLoan.Borrowers = append(recLoan.Borrowers, Borrower{v, try[v]})
	}
	return recLoan

}

func RecordJSON(record LoanRecord, path string) error {
	var jsonData, err = json.Marshal(record)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonData)
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		panic(err)
	}
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
}
