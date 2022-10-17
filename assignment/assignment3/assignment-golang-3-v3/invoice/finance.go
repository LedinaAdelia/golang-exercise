package invoice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if fi.Date == "" {
		return InvoiceData{}, EmptyDateErr
	}
	if fi.Status == "" && fi.Status != PAID && fi.Status != UNPAID {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}
	value, err := calculateRecord(fi.Details)
	if err != nil {
		return InvoiceData{}, err
	}
	formatDate := DateFormat(fi.Date)
	invoiceRecord := InvoiceData{
		Date:         formatDate,
		Departemen:   Finance,
		TotalInvoice: value,
	}
	return invoiceRecord, nil // TODO: replace this
}
func calculateRecord(data []Detail) (float64, error) {
	if len(data) == 0 {
		return 0.0, errors.New("invoice details is empty")
	}

	temp := 0.0
	for _, v := range data {
		temp += float64(v.Total)
	}

	if temp <= 0 {
		return 0.0, errors.New("total price is not valid")
	}
	return temp, nil
}

func DateFormat(date string) string {
	a := strings.Split(date, "/")
	month, _ := strconv.Atoi(a[1])
	m := time.Month(month)

	res := fmt.Sprintf("%s-%s-%s", a[0], m, a[2])
	return res
}
