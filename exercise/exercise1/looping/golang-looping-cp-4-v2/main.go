package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	r := []rune(email)
	checkDomain := false
	checkTLD := false
	domain := ""
	TLD := ""
	for _, v := range r {
		if string(v) == "@" {
			checkDomain = true
			checkTLD = false
			continue
		} else if string(v) == "." {
			checkTLD = true
			checkDomain = false
		}

		if checkDomain == true {
			domain += string(v)
		} else if checkTLD == true {
			TLD += string(v)
		}
	}
	formattedTLD := strings.Replace(TLD, ".", "", 1)
	result := fmt.Sprintf("Domain: %s dan TLD: %s", domain, formattedTLD)
	return result // TODO: replace this,

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
