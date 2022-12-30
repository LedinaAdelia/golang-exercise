package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	checkDomain := false
	checkTLD := false
	domain := ""
	TLD := ""
	for _, v := range email {
		if string(v) == "@" {
			checkDomain = true
			checkTLD = false
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
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, strings.Replace(TLD, ".", "", 1))

}

func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
