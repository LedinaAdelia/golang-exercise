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

	final := []invoice.InvoiceData{}
	temp := map[string]float64{}
	a := invoice.InvoiceData{}
	for _, v := range output {
		if v.Departemen == invoice.Finance {
			temp[v.Date] += v.TotalInvoice
			a = invoice.InvoiceData{
				Date:         v.Date,
				Departemen:   v.Departemen,
				TotalInvoice: temp[v.Date],
			}
			final = append(final, a)
		} else if v.Departemen == invoice.Marketing {
			temp[v.Date] += v.TotalInvoice
			a = invoice.InvoiceData{
				Date:         v.Date,
				Departemen:   v.Departemen,
				TotalInvoice: temp[v.Date],
			}
			final = append(final, a)
		}
	}
	final = append(final, a)
	return final, nil // TODO: replace this
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date: "01/02/2020",
			Details: []invoice.Detail{
				{Description: "pembelian nota", Total: 4000},
				{Description: "Pembelian alat tulis", Total: 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date: "01/02/2020",
			Details: []invoice.Detail{
				{Description: "pembelian nota", Total: 4000},
				{Description: "Pembelian alat tulis", Total: 3000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "03-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.WarehouseInvoice{
			Date: "03-February-2020",
			Products: []invoice.Product{
				{Name: "product A", Unit: 10, Price: 10000, Discount: 0.1},
				{Name: "product C", Unit: 5, Price: 15000, Discount: 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "04/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
		invoice.MarketingInvoice{
			Date:        "04/02/2020",
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
