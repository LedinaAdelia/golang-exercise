package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	temp := map[string][]string{}
	for _, v := range transactions {
		if v.Type == "income" {
			str_amount := strconv.Itoa(v.Amount)
			try := fmt.Sprintf("income:%s", str_amount)
			temp[v.Date] = append(temp[v.Date], try)
		} else if v.Type == "expense" {
			str_amount := strconv.Itoa(v.Amount)
			try := fmt.Sprintf("expense:%s", str_amount)
			temp[v.Date] = append(temp[v.Date], try)
		}
	}
	count := countDate(temp)
	sort.Strings(count)
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)
	for _, data := range count {
		if len(count) == 1 {
			_, _ = datawriter.WriteString(data + "")
		} else if count[len(count)-1] == data {
			_, _ = datawriter.WriteString(data + "")
		} else {
			_, _ = datawriter.WriteString(data + "\n")
		}
	}
	datawriter.Flush()
	file.Close()

	return nil
}

func countDate(data map[string][]string) []string {
	newMap := map[string]int{}
	fix := []string{}
	for key, a := range data {
		for _, b := range a {
			try := strings.Split(b, ":")
			if try[0] == "income" {
				conv, _ := strconv.Atoi(try[1])
				newMap[key] += conv
			} else if try[0] == "expense" {
				conv, _ := strconv.Atoi(try[1])
				newMap[key] -= conv
			}
		}
	}

	for key, v := range newMap {
		if v >= 0 {
			res := fmt.Sprintf("%s;income;%d", key, v)
			fix = append(fix, res)
		} else if v <= 0 {
			convTotal := math.Abs(float64(v))
			res := fmt.Sprintf("%s;expense;%.0f", key, convTotal)
			fix = append(fix, res)
		}
	}
	return fix
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		// {"01/01/2021", "income", 1000},
		// {"02/01/2021", "income", 20000},
		// {"02/01/2021", "income", 233000},
		// {"03/01/2021", "income", 20000},
		// {"03/01/2021", "income", 22000},
		// {"04/01/2021", "income", 24000},
		// {"04/01/2021", "income", 20000},
		// {"05/01/2021", "income", 20000},
		// {"05/01/2021", "income", 50000},
		// {"06/01/2021", "income", 60000},
		// {"06/01/2021", "income", 70000},
		// {"07/01/2021", "income", 80000},
		// {"07/01/2021", "income", 110000},
		// {"07/01/2021", "income", 120000},
		// {"07/01/2021", "income", 111200},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
