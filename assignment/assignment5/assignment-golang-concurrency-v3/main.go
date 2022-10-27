package main

import (
	"fmt"
	"strings"
	"sync"
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
	}
	ch <- website
}

var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)
	var err error
	var wg sync.WaitGroup
	a := []RowData{}
	for _, website := range data {
		wg.Add(1)
		go func(website RowData) {
			go FuncProcessGetTLD(website, ch, errCh)
			fmt.Println(<-ch)
			a = append(a, <-ch)
			fmt.Println("len", a)
			wg.Done()
		}(website)
	}
	wg.Wait()

	close(ch)
	close(errCh)
	return a, err
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
		{RankWebsite: 1, Domain: "google.com", Valid: true, RefIPs: 2404064},
		{RankWebsite: 2, Domain: "facebook.com", Valid: true, RefIPs: 2547862},
		{RankWebsite: 3, Domain: "della.com", Valid: true, RefIPs: 2067945},
		{RankWebsite: 4, Domain: "", Valid: true, RefIPs: 2046616},
		{RankWebsite: 5, Domain: "instagram.com", Valid: true, RefIPs: 1694854},
	}
	fmt.Println(FilterAndFillData(".com", a))
}
