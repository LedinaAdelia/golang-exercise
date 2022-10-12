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
	fmt.Println(count)
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
	return nil // TODO: replace this
}

func FormatRupiah(number int) string {
	return ""
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	GetDayDifference(dateRange)
	return nil
}

func main() {
	res := GetSalaryOverview("25 February - 5 March 2020", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}
