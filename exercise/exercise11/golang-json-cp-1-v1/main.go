package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

type Study struct {
	Study_name   string  `json:"study_name"`
	Study_credit float64 `json:"study_credit"`
	Grade        string  `json:"grade"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	file, _ := ioutil.ReadFile(filename)
	data := Report{}
	_ = json.Unmarshal([]byte(file), &data)
	return data, nil
}

func GradePoint(report Report) float64 {
	totalNilai := 0.0
	totalKredit := 0.0
	fix := 0.0

	if len(report.Studies) == 0 {
		return 0.0
	}

	for _, v := range report.Studies {
		if v.Grade == "A" {
			totalNilai += 4 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "AB" {
			totalNilai += 3.5 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "B" {
			totalNilai += 3 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "BC" {
			totalNilai += 2.5 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "C" {
			totalNilai += 2 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "CD" {
			totalNilai += 1.5 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "D" {
			totalNilai += 1 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "DE" {
			totalNilai += 0.5 * v.Study_credit
			totalKredit += v.Study_credit
		} else if v.Grade == "E" {
			totalNilai += 0 * v.Study_credit
			totalKredit += v.Study_credit
		}
	}
	total := totalNilai / totalKredit
	formatTotal := fmt.Sprintf("%.2f", total)

	fix, _ = strconv.ParseFloat(formatTotal, 64)
	return fix // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
