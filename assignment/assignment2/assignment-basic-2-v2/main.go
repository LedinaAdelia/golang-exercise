package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetDayDifference(date string) int {
	s := strings.Split(date, " ")
	checkDay1, _ := strconv.Atoi(s[0])
	checkDay2, _ := strconv.Atoi(s[3])
	if checkDay1 < 10 {
		s[0] = fmt.Sprintf("0%d", checkDay1)
	} else if checkDay2 < 10 {
		s[3] = fmt.Sprintf("0%d", checkDay2)
	}
	count := 1
	if len(s) != 6 {
		fmt.Println("Format Tanggal Salah")
	} else {
		formatDay1 := fmt.Sprintf("%s-%s-%s", s[5], s[1], s[0])
		formatDay2 := fmt.Sprintf("%s-%s-%s", s[5], s[4], s[3])
		firstDay := changeFormatDate(formatDay1)
		lastDay := changeFormatDate(formatDay2)
		a := strings.Split(firstDay, "-")
		b := strings.Split(lastDay, "-")
		date, _ := strconv.Atoi(a[0])
		month, _ := strconv.Atoi(a[1])
		year, _ := strconv.Atoi(a[2])
		date2, _ := strconv.Atoi(b[0])
		month2, _ := strconv.Atoi(b[1])
		year2, _ := strconv.Atoi(b[2])
		d := Date(year, month, date)
		d2 := Date(year2, month2, date2)
		days := d2.Sub(d).Hours() / 24
		if (month%3) == 0 || month == 1 {
			days -= 1
		}
		count += int(days)
	}
	return count
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func changeFormatDate(d string) string {
	layout := "2006-January-02"
	parse, _ := time.Parse(layout, d)
	date := parse.Format("02-01-2006")
	return date
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	salary := 50000
	temp := map[string]string{}
	count := map[string]int{}
	for i, a := range data {
		for _, b := range a {
			if i < rangeDay {
				count[b] += salary
				con := FormatRupiah(count[b])
				temp[b] = con
			}
		}
	}
	return temp // TODO: replace this
}

func FormatRupiah(number int) string {
	arrChange := strconv.Itoa(number)
	splitChange := strings.Split(arrChange, "")
	cResultReverse := ""
	cResult := ""
	count := 0
	for i := 0; i < len(splitChange); i++ {
		if count == 2 && i != (len(splitChange)-1) {
			count = 0
			cResultReverse += splitChange[len(splitChange)-i-1] + "."

		} else {
			count++
			cResultReverse += splitChange[len(splitChange)-i-1]
		}
	}
	result := strings.Split(cResultReverse, "")
	for j := 0; j < len(result); j++ {
		cResult += result[len(result)-j-1]
	}

	return fmt.Sprintf("Rp %s", cResult)
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	rangeDay := GetDayDifference(dateRange)
	countSalary := GetSalary(rangeDay, data)
	return countSalary
}

func main() {
	res := GetSalaryOverview("21 February - 22 February 2021", [][]string{
		{"a", "b"},
		{"a", "c"},
	})
	fmt.Println(res)
}
