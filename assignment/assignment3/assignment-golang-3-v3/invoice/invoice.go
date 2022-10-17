package invoice

import "errors"

type InvoiceData struct {
	Date         string
	TotalInvoice float64
	Departemen   DepartmentName
}

type DepartmentName string

const (
	Finance   DepartmentName = "finance"
	Warehouse DepartmentName = "warehouse"
	Marketing DepartmentName = "marketing"
)

var (
	EmptyDateErr = errors.New("invoice date is empty")
)

type Invoice interface {
	RecordInvoice() (InvoiceData, error)
}
