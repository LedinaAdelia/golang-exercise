package invoice

import (
	"errors"
	"fmt"
)

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if wi.Date == "" {
		return InvoiceData{}, EmptyDateErr
	}
	if wi.InvoiceType == "" && wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {
		return InvoiceData{}, errors.New("invoice type is not valid")
	}
	value, err := countTotal(wi.Products)
	if err != nil {
		return InvoiceData{}, err
	}
	fmt.Println("")
	invoiceRecord := InvoiceData{
		Date:         wi.Date,
		Departemen:   Warehouse,
		TotalInvoice: value,
	}
	return invoiceRecord, nil // TODO: replace this
}

func countTotal(data []Product) (float64, error) {
	if len(data) == 0 {
		return 0.0, errors.New("invoice products is empty")
	}
	temp := map[string]float64{}
	total := 0.0
	for _, v := range data {
		if v.Unit == 0 {
			return 0.0, errors.New("unit product is not valid")
		} else {
			tempTotal := (float64(v.Unit) * v.Price)
			temp[v.Name] = tempTotal - (v.Discount * tempTotal)
		}
	}
	for _, v := range temp {
		total += v
	}
	if total <= 0 {
		return 0.0, errors.New("price product is not valid")
	}
	return total, nil
}
