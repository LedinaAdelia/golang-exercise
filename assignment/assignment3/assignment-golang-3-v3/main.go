package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)


func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	return nil, nil // TODO: replace this
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
				{"product A", 10, 10000, 0.1},
				{"product C", 5, 15000, 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "01/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
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
