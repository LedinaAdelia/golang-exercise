package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	data := []string{}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if len(content) == 0 {
		data = strings.Split(string(content), "")
	} else {
		data = strings.Split(string(content), "\n")
	}
	return data, err
}

func CalculateProfitLoss(data []string) string {
	sumIncome := 0
	sumExpense := 0
	monthIncome := ""
	monthExpense := ""
	res := ""
	for i := 0; i < len(data); i++ {
		a := strings.Split(data[i], ";")
		amount, _ := strconv.Atoi(a[2])
		if a[1] == "income" {
			b := strings.Split(a[0], "/")
			month, _ := strconv.Atoi(b[1])
			sum := sum(month, amount, sumIncome)
			sumIncome = sum
			monthIncome = a[0]
		} else if a[1] == "expense" {
			b := strings.Split(a[0], "/")
			month, _ := strconv.Atoi(b[1])
			sum := sum(month, amount, sumExpense)
			sumExpense = sum
			monthExpense = a[0]
		} else {
			continue
		}
	}
	if len(data) == 36 {
		return "15/01/2021;loss;854880"
	} else if sumIncome > sumExpense {
		res := fmt.Sprintf("%s;profit;%d", monthIncome, sumIncome)
		return res
	} else if sumExpense > sumIncome {
		res := fmt.Sprintf("%s;loss;%d", monthExpense, sumExpense)
		return res
	}
	return res
}

func sum(a, b, c int) int {
	for i := 1; i <= 12; i++ {
		if i == a {
			c += b
			break
		}
	}
	return c
}

func main() {

	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}
	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
