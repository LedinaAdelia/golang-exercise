package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
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
	// for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
	// 	r[i], r[j] = byte(fmt.Sprintf("_%s".r[j])), r[i]
	// }

	return char
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
