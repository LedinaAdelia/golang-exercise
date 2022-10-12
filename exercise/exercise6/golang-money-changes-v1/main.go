package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	// nggak rusak see, ternyata kepencet HEHEHEHEHEHEHEH :) <3
	//apanya yag rusak???
	//apamyaa????
	// INI
	// itu diatas
	// kepencet capslock, oalaaa iya dekyanggg
	//
	//okay, KOKKK CAPSLOCK
	counts := 0
	sum := 0
	for _, a := range products {
		counts += a.Price + a.Tax
	}
	sum = amount - counts
	hm := []int{}
	count := false
	for (sum >= 0) && !count {
		if sum >= 1000 {
			sum -= 1000
			hm = append(hm, 1000)
		} else if (sum >= 500) && (sum < 1000) {
			sum -= 500
			hm = append(hm, 500)
		} else if (sum >= 200) && (sum < 500) && (sum < 1000) {
			sum -= 200
			hm = append(hm, 200)
		} else if (sum >= 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 100
			hm = append(hm, 100)
		} else if (sum >= 50) && (sum < 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 50
			hm = append(hm, 50)
		} else if (sum >= 20) && (sum < 50) && (sum < 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 20
			hm = append(hm, 20)
		} else if (sum >= 10) && (sum < 20) && (sum < 50) && (sum < 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 10
			hm = append(hm, 10)
		} else if (sum >= 5) && (sum < 10) && (sum < 20) && (sum < 50) && (sum < 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 5
			hm = append(hm, 5)
		} else if (sum >= 1) && (sum < 5) && (sum < 10) && (sum < 20) && (sum < 50) && (sum < 100) && (sum < 200) && (sum < 500) && (sum < 1000) {
			sum -= 1
			hm = append(hm, 1)
		} else {
			count = true
		}
	}
	return hm
}

func main() {
	cart := []Product{
		Product{
			Name:  "Baju",
			Price: 20000,
			Tax:   500,
		},
		Product{
			Name:  "Celana",
			Price: 20000,
			Tax:   500,
		},
	}
	fmt.Println(cart)
	fmt.Println(MoneyChanges(100000, cart))
}
