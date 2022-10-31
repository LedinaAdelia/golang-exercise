package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	// buat variabel buat string split si string dari calculate
	// split := .......
	split := strings.Split(calculate, " ")
	// di package internal funcnya berparameter struct base jadi buat variabel basenya
	// variabel base itu yang paling depan/awal (jadi yang split[0] harusnya)
	// variabelnya masih berubah string rubah ke float32
	// base, _ := ......
	base, _ := strconv.ParseFloat(split[0], 64)
	// Panggil function di package internal caranya kayak gini
	// Cal := internal.NewCalculator(base)
	Cal := internal.NewCalculator(float32(base))
	// buat looping sepanjang dari array split (pake looping statement aja)
	// for{
	// buat condition dimana kalo misal di split ada (*,/,+,-)
	// if split[i]=="*"{
	//rubah angka setelah string index itu ke float
	// kalo misal ada 2*4 (2 itu base, setelah * kan angka 4) nah 4 dirubah ke float
	// num, _:= ......
	// panggil function multiply dari package internal
	// Cal.Multiply(float32(num))
	// }
	// dst.
	//}
	for i := 0; i < len(split); i++ {
		if split[i] == "*" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Multiply(float32(num))
		} else if split[i] == "/" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Divide(float32(num))
		} else if split[i] == "+" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Add(float32(num))
		} else if split[i] == "-" {
			num, _ := strconv.ParseFloat(split[i+1], 64)
			Cal.Subtract(float32(num))
		}
	}
	return Cal.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("100 / 2 / 5 + 10 - 5")
	fmt.Println(res)
}
