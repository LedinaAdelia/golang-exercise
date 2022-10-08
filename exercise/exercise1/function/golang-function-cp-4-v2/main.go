package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	arrData := []string(data)
	result := ""
	for i := 0; i < len(arrData); i++ {
		if strings.Contains(arrData[i], input) {
			result += arrData[i] + ","
		}
	}
	s1 := result
	if last := len(s1) - 1; last >= 0 && s1[last] == ',' {
		s1 = s1[:last]
	}
	return s1 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
	fmt.Println(FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
}
