package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	if (math+science+english+indonesia)/4 == 100 {
		return "Sempurna"
	} else if (math+science+english+indonesia)/4 >= 90 {
		return "Sangat Baik"
	} else if (math+science+english+indonesia)/4 >= 80 {
		return "Baik"
	} else if (math+science+english+indonesia)/4 >= 70 {
		return "Cukup"
	} else if (math+science+english+indonesia)/4 > 60 {
		return "Kurang"
	} else if (math+science+english+indonesia)/4 <= 60 {
		return "Sangat kurang"
	}
	return "" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
