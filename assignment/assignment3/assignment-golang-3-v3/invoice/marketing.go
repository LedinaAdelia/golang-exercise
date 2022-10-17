package invoice

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	if mi.Date == "" {
		return InvoiceData{}, EmptyDateErr
	}
	formatDate := DateFormat(mi.Date)

	totalDay, err := countDay(mi.StartDate, mi.EndDate)
	if err != nil {
		return InvoiceData{}, err
	}
	totalPrice, err := countPrice(totalDay, mi.PricePerDay, mi.AnotherFee)
	if err != nil {
		return InvoiceData{}, err
	}
	invoiceRecord := InvoiceData{
		Date:         formatDate,
		Departemen:   Marketing,
		TotalInvoice: totalPrice,
	}
	return invoiceRecord, nil
}

func countDay(start, end string) (float64, error) {
	if start == "" || end == "" {
		return 0.0, errors.New("travel date is empty")
	}
	a := strings.Split(start, "/")
	b := strings.Split(end, "/")
	date, _ := strconv.Atoi(a[0])
	month, _ := strconv.Atoi(a[1])
	year, _ := strconv.Atoi(a[2])
	date2, _ := strconv.Atoi(b[0])
	month2, _ := strconv.Atoi(b[1])
	year2, _ := strconv.Atoi(b[2])
	d := Date(year, month, date)
	d2 := Date(year2, month2, date2)
	days := d2.Sub(d).Hours() / 24
	return days + 1, nil
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func countPrice(day float64, roadprice, fee int) (float64, error) {
	total := (day * float64(roadprice)) + float64(fee)
	if total <= 0 {
		return 0.0, errors.New("price per day is not valid")
	}
	return total, nil
}
