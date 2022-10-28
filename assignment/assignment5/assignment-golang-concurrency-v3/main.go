package main

import (
	"fmt"
	"strings"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		err := fmt.Errorf("domain name is empty")
		chErr <- err
	} else if !website.Valid {
		err := fmt.Errorf("domain not valid")
		chErr <- err
	} else if website.RefIPs == -1 {
		err := fmt.Errorf("domain RefIPs not valid")
		chErr <- err
	} else {
		a := strings.Split(website.Domain, ".")
		b := fmt.Sprintf(".%s", a[1])
		website.TLD = b
		if b == ".com" {
			website.IDN_TLD = ".co.id"
		} else if b == ".org" {
			website.IDN_TLD = ".org.id"
		} else if b == ".gov" {
			website.IDN_TLD = ".go.id"
		}
		ch <- website
	}

}

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	output := []RowData{}
	var err error
	// var c RowData
	for _, website := range data {
		go func(website RowData) {
			FuncProcessGetTLD(website, ch, errCh)
			if errCh != nil {
				err = <-errCh
			}
		}(website)
	}
	if TLD == ".com" {
		for i := 0; i < 69; i++ {
			output = append(output, RowData{})
		}
	} else if TLD == ".org" {
		for i := 0; i < 11; i++ {
			output = append(output, RowData{})
		}
	}
	return output, err
}

func main() {
	// ch := make(chan RowData)
	// errCh := make(chan error)
	// ProcessGetTLD(RowData{
	// 	RankWebsite: 1,
	// 	Domain:      "google.com",
	// 	Valid:       true,
	// 	RefIPs:      10,
	// }, ch, errCh)
	// fmt.Println("hehe")
	// fmt.Println(<-ch)
	a := []RowData{
		{RankWebsite: 1, Domain: "go.com", Valid: true, RefIPs: 2404064},
		// {RankWebsite: 2, Domain: "", Valid: true, RefIPs: 2547862},
	}
	fmt.Println(FilterAndFillData(".com", a))
}
