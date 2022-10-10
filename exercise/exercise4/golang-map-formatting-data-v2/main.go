package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	a := make(map[string][]string)
	for i := 0; i < len(data); i++ {
		hm := strings.Split(data[i], "-")
		// fmt.Println(hm)
		a[hm[0]] = append(a[hm[0]], hm[3])
	}
	output := make(map[string][]string)
	for key, value := range a {
		for i := 0; i < len(a[key]); i++ {
			if key == "phone" {
				output[key] = append(output[key], value[i])
			} else if i%2 == 0 {
				b := value[i] + " " + value[i+1]
				output[key] = append(output[key], b)
			}
		}
	}
	return output // TODO: replace this
}

func main() {
	a := []string{
		"account-0-first-John",
		"account-0-last-Doe",
		"account-1-first-Jane",
		"account-1-last-Doe",
		"address-0-first-Jaksel",
		"address-0-last-Jakarta",
		"address-1-first-Bandung",
		"address-1-last-Jabar",
		"phone-0-first-081234567890",
		"phone-1-first-081234567891",
	}
	fmt.Println(ChangeOutput(a))
}
