package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	pairs := [][]string{}
	for key, value := range mapData {
		pairs = append(pairs, []string{key, value})
	}
	return pairs // TODO: replace this
}

func main() {
	a := map[string]string{}
	fmt.Println(MapToSlice(a))
}
