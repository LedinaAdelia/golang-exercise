package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	a := make(map[string]float32)
	for i := 0; i < len(data); i++ {
		c := strings.Split(data[i], ":")
		b := c[0] + "-" + c[1]
		hm := try(c[3], day)
		fmt.Println(hm)
		if hm == true {
			if day == "senin" || day == "rabu" || day == "jumat" {
				try, _ := strconv.ParseFloat(c[2], 64)
				price := try + (try * 0.1)
				a[b] = float32(price)
			} else if day == "selasa" || day == "kamis" || day == "sabtu" {
				try, _ := strconv.ParseFloat(c[2], 64)
				price := try + (try * 0.05)
				a[b] = float32(price)
			}
		} else {
			fmt.Println("salah")
		}
		fmt.Println("")
	}
	return a // TODO: replace this
}

func try(a string, b string) bool {
	JKT := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	BDG := []string{"rabu", "kamis", "sabtu"}
	BKS := []string{"selasa", "kamis", "jumat"}
	DPK := []string{"senin", "selasa"}
	hai := false
	fmt.Println(b)
	if a == "JKT" {
		for _, v := range JKT {
			if b == v {
				hai = true
			}
		}
	} else if a == "BDG" {
		for _, v := range BDG {
			if b == v {
				hai = true
			}
		}
	} else if a == "BKS" {
		for _, v := range BKS {
			if b == v {
				hai = true
			}
		}
	} else if a == "DPK" {
		for _, v := range DPK {
			if b == v {
				hai = true
			}
		}
	}
	return hai
}

func main() {
	data := []string{
		"Anggi:Anggraini:20000:DPK",
		"Andi:Sukirman:15000:JKT",
	}
	day := "selasa"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
