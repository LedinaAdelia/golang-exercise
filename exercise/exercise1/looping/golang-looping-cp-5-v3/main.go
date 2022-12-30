package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	arrString := strings.Split(str, " ")
	var reversedArrString string
	for i := 0; i < len(arrString); i++ {
		if i == len(arrString)-1 {
			reversedArrString += reverseString(arrString[i])
		} else {
			reversedArrString += reverseString(arrString[i]) + " "
		}
	}
	return reversedArrString

}
func reverseString(str string) string {
	var char string
	for i := 0; i < len(str); i++ {
		if i == 0 {
			if unicode.IsUpper(rune(str[i])) {
				char += strings.ToUpper(string(str[len(str)-i-1]))
			} else {
				char += strings.ToLower(string(str[len(str)-i-1]))
			}
		} else {
			char += strings.ToLower(string(str[len(str)-i-1]))
		}
	}
	return char
}

func main() {
	fmt.Println(ReverseWord("Aku Sayang Mama"))
}
