package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	arrString := strings.Split(str, " ")
	var reversedArrString []string
	for i := 0; i < len(arrString); i++ {
		reversedArrString = append(reversedArrString, reverseString(arrString[i]))
	}
	return strings.Join(reversedArrString, " ") // TODO: replace this

}
func reverseString(str string) string {
	r := []rune(str)
	var char string
	for i := 0; i < len(r); i++ {
		if i == 0 {
			if unicode.IsUpper(r[i]) == true {
				char += strings.ToUpper(string(r[len(r)-i-1]))
			} else {
				char += strings.ToLower(string(r[len(r)-i-1]))
			}
		} else {
			char += strings.ToLower(string(r[len(r)-i-1]))
		}
	}
	return char
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Mama"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
	fmt.Println(ReverseWord("ini terlalu mudah"))
	fmt.Println(ReverseWord("KITA SELALU BERSAMA"))
}
