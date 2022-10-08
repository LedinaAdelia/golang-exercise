package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindShortestName(names string) string {
	arrString := strings.Split(names, " ")
	if strings.Contains(names, ";") {
		arrString = strings.Split(names, ";")
	} else if strings.Contains(names, ",") {
		arrString = strings.Split(names, ",")
	}
	sort.Strings(arrString)
	fmt.Println(arrString)
	minValue := len(arrString[0])
	minValuePair := arrString[0]
	fmt.Println(minValuePair)
	for i := 1; i < len(arrString); i++ {
		number := len(arrString[i])
		fmt.Println(number)
		fmt.Println(arrString[i])
		if minValue > number {
			minValue = number
			minValuePair = arrString[i]
		}
	}
	return minValuePair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan"))
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))
	fmt.Println(FindShortestName("A,B,C,D,E"))
	fmt.Println(FindShortestName("Hanif;Joko;Indah;Ari;Intan"))
}
