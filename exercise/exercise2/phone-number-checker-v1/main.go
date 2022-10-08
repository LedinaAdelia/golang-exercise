package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	arrNumber := strings.Split(number, "")
	a := arrNumber[0] + arrNumber[1] + arrNumber[2]
	b := arrNumber[0] + arrNumber[1]
	if (a == "628") && (len(arrNumber) >= 11) {
		c := arrNumber[3] + arrNumber[4]
		var cRes string
		findProvider(c, &cRes)
		*result = cRes
	} else if (b == "08") && (len(arrNumber) >= 10) {
		c := arrNumber[2] + arrNumber[3]
		var cRes string
		findProvider(c, &cRes)
		*result = cRes
	} else {
		*result = "invalid"
	}
}

func findProvider(try string, rest *string) {
	arrTelkom := []string{"11", "12", "13", "14", "15"}
	arrIndosat := []string{"16", "17", "18", "19"}
	arrXl := []string{"21", "22", "23"}
	arrTri := []string{"27", "28", "29"}
	arrAs := []string{"52", "53"}
	arrSf := []string{"81", "82", "83", "84", "85", "86", "87", "88"}
	a := false
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("hai")
	// 	if try == arrTelkom[i] {
	// 		*rest = "Telkomsel"
	// 		a = true

	// 	} else if try == arrIndosat[i] {
	// 		*rest = "Indosat"
	// 		a = true

	// 	} else if try == arrXl[i] {
	// 		*rest = "Xl"
	// 		a = true

	// 	} else if try == arrTri[i] {
	// 		*rest = "Tri"
	// 		a = true

	// 	} else if try == arrAs[i] {
	// 		*rest = "As"
	// 		a = true

	// 	} else if try == arrSf[i] {
	// 		*rest = "Smartfren"
	// 		a = true
	// 	} else {
	// 		a = false
	// 	}
	// }
	for _, v := range arrTelkom {
		if try == v {
			*rest = "Telkomsel"
			a = true
		}
	}
	for _, v := range arrIndosat {
		if try == v {
			*rest = "Indosat"
			a = true
		}
	}
	for _, v := range arrAs {
		if try == v {
			*rest = "AS"
			a = true
		}
	}
	for _, v := range arrXl {
		if try == v {
			*rest = "XL"
			a = true
		}
	}
	for _, v := range arrTri {
		if try == v {
			*rest = "Tri"
			a = true
		}
	}
	for _, v := range arrSf {
		if try == v {
			*rest = "Smartfren"
			a = true
		}
	}
	if a == false {
		*rest = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081434561234"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
