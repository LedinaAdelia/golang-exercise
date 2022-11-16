package main

import (
	"fmt"
)

func ReverseString(str string) string {
	r := []rune(str)
	var char string
	for i := 0; i < len(r); i++ {
		if i == 0 || string(r[len(r)-i-1]) == " " || string(r[len(r)-i]) == " " {
			char += string(r[len(r)-i-1])
		} else {
			char += fmt.Sprintf("_%s", string(r[len(r)-i-1]))
		}
	}
	return char
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
