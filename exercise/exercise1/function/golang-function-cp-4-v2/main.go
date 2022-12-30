package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	res := ""
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i], input) {
			res += data[i] + ","
		}
	}
	if last := len(res) - 1; last >= 0 && res[last] == ',' {
		res = res[:last]
	}
	return res
}

func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
