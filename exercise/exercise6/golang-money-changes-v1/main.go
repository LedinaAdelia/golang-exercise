package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	resultChange := []int{}
	total := 0
	for _, v := range products {
		total += v.Price + v.Tax
	}
	change := amount - total
	check := false
	for change >= 0 && !check {
		if change >= 1000 {
			change -= 1000
			resultChange = append(resultChange, 1000)
		} else if change >= 500 {
			change -= 500
			resultChange = append(resultChange, 500)
		} else if change >= 200 {
			change -= 200
			resultChange = append(resultChange, 200)
		} else if change >= 100 {
			change -= 100
			resultChange = append(resultChange, 100)
		} else if change >= 50 {
			change -= 50
			resultChange = append(resultChange, 50)
		} else if change >= 20 {
			change -= 20
			resultChange = append(resultChange, 20)
		} else if change >= 10 {
			change -= 10
			resultChange = append(resultChange, 10)
		} else if change >= 5 {
			change -= 5
			resultChange = append(resultChange, 5)
		} else if change >= 1 {
			change -= 1
			resultChange = append(resultChange, 1)
		} else {
			check = true
		}
	}

	return resultChange
}

func main() {
	cart := []Product{
		{
			Name:  "Baju",
			Price: 5000,
			Tax:   500,
		},
		{
			Name:  "Celana",
			Price: 3000,
			Tax:   300,
		},
	}

	fmt.Println(MoneyChanges(10000, cart))
}
