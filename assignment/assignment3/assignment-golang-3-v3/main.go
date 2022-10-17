package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	output := []invoice.InvoiceData{}
	for _, v := range data {
		invoiceRecord, err := v.RecordInvoice()
		if err != nil {
			return nil, err
		}
		output = append(output, invoiceRecord)
	}
	temp := map[string]float64{}
	try := invoice.InvoiceData{}
	result := []invoice.InvoiceData{}
	for _, v := range output {
		if v.Departemen == "finance" {
			temp[v.Date] += v.TotalInvoice
			try.Date = v.Date
			try.Departemen = v.Departemen
			try.TotalInvoice = temp[v.Date]
		}
	}
	result = append(result, try)
	return result, nil // TODO: replace this
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "01-February-2020",
			Products: []invoice.Product{
				{"product A", 2, 10000, 0.1},
				{"product C", 3, 20000, 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "01/01/2022",
			EndDate:     "03/01/2022",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
