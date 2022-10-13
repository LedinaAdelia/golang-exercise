package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	var final = []map[string]interface{}{}
	for _, a := range data {
		s := strings.Split(a, ";")
		fmt.Println(len(s))
		age, _ := strconv.Atoi(s[1])
		height, _ := strconv.ParseFloat(s[3], 64)
		isMarried, _ := strconv.ParseBool(s[4])
		temp := map[string]interface{}{}
		if s[3] != "" && s[4] != "" {
			temp = map[string]interface{}{
				"name":      s[0],
				"age":       age,
				"address":   s[2],
				"height":    height,
				"isMarried": isMarried,
			}
		} else if s[3] == "" && s[4] == "" {
			temp = map[string]interface{}{
				"name":    s[0],
				"age":     age,
				"address": s[2],
			}
		} else if s[3] == "" {
			temp = map[string]interface{}{
				"name":      s[0],
				"age":       age,
				"address":   s[2],
				"isMarried": isMarried,
			}
		} else if s[4] == "" {
			temp = map[string]interface{}{
				"name":    s[0],
				"age":     age,
				"address": s[2],
				"height":  height,
			}
		}
		final = append(final, temp)
	}

	return final // TODO: replace this
}

func main() {
	data := []string{
		"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;;false",
		"Santi;27;Jakarta;160;",
		"Rudi;23;Jakarta;161.1;",
		"Rudi;23;Jakarta;166.5;false",
		"Rudi;23;Jakarta;;",
	}
	fmt.Println(PopulationData(data))
}
