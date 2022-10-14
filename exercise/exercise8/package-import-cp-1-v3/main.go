package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	a := strings.Split(calculate, " ")
	var res float32
	base, _ := strconv.ParseFloat(a[0], 64)
	try := internal.Calculator{float32(base)}
	for i := 0; i < len(a); i++ {
		if a[i] == "*" {
			num, _ := strconv.ParseFloat(a[i+1], 64)
			try.Multiply(float32(num))
		} else if a[i] == "/" {
			num, _ := strconv.ParseFloat(a[i+1], 64)
			try.Divide(float32(num))
		} else if a[i] == "+" {
			num, _ := strconv.ParseFloat(a[i+1], 64)
			try.Add(float32(num))
		} else if a[i] == "-" {
			num, _ := strconv.ParseFloat(a[i+1], 64)
			try.Subtract(float32(num))
		}
	}
	res = try.Result()
	return res // TODO: replace this
}

func main() {
	res := AdvanceCalculator("100 / 2 / 5 + 10 - 5")
	fmt.Println(res)
}
