package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Reverse(str string) string {
	arrReverse := strings.Split(str, "")
	rReverse := ""
	for i := 0; i < len(arrReverse); i++ {
		rReverse += arrReverse[(len(arrReverse) - i - 1)]
	}
	return rReverse // TODO: replace this
}

func Generate(str string) string {
	str = Reverse(str)
	arrGenerate := strings.Split(str, "")
	vokal := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	output := ""
	for i := 0; i < len(arrGenerate); i++ {
		a := arrGenerate[i]
		for j := 0; j < len(vokal); j++ {
			fmt.Println("sebelum: ", a)
			if a == vokal[len(vokal)-1] {
				a = vokal[5]
				fmt.Print("sesudah: ", a)
				break
			} else if a == vokal[j] {
				a = vokal[j+1]
				break
			}
		}

		if arrGenerate[i] == " " {
			output += ""
		} else if strings.ToUpper(a) == a {
			fmt.Println("sebelum: ", strings.ToLower(a))
			output += strings.ToLower(a)
			fmt.Println("Upper: ", output)
		} else {
			fmt.Println("sebelum-2: ", strings.ToLower(a))
			output += strings.ToUpper(a)
		}
	}
	return output // TODO: replace this
}

func switchTextCase(v string) string {
	if strings.ToUpper(v) == v {
		return strings.ToLower(v)
	}
	return strings.ToUpper(v)
}

func CheckPassword(str string) string {
	arrString := strings.Split(str, "")
	word, _ := regexp.Compile(`[A-Za-z1-9]`)
	char, _ := regexp.Compile(`[~!@#$%^&*]`)
	len1 := len(arrString) >= 7
	len2 := len(arrString) >= 14
	conWord := word.MatchString(str)
	conChar := char.MatchString(str)

	if (len2 == true) && (conWord == true) && (conChar == true) {
		return "kuat"
	} else if (len1 == true) && (conWord == true) && (conChar == true) {
		return "sedang"
	} else if (len1 == true) && (conWord == true) {
		return "lemah"
	} else if len1 == false {
		return "sangat lemah"
	}
	return "" // TODO: replace this
}

func PasswordGenerator(base string) (string, string) {
	res := Generate(base)
	fmt.Println(len(res))
	check := CheckPassword(res)
	return res, check // TODO: replace this
}

func main() {
	data := "Seman-8gat Pagi!!" // bisa digunakan untuk melakukan debug
	try := "HAIU"
	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)

	res2, check2 := PasswordGenerator(try)
	fmt.Println(res2)
	fmt.Println(check2)
}
