package main

import (
	"fmt"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	for i := 0; i < len(transactions); i++ {
		fmt.Println(transactions[i])
	}
	// return errors.New("not implemented") // TODO: replace this
	return nil
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
