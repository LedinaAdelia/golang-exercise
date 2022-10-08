package main

import (
	"fmt"
)

func SlurredTalk(words *string) {

}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Steven is me"
	SlurredTalk(&words)
	fmt.Println(words)
}
