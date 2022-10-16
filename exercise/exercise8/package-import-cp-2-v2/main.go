package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)
	return service
}

// gunakan untuk debugging
func main() {
	database := database.NewDatabase()
	service := CashierApp(database)

	service.AddCart("Kaos Polos", 2)
	service.AddCart("Kaos sablon", 1)
	service.AddCart("Celana hitam", 3)
	fmt.Println(service.RemoveCart("Celana hitam"))
	paymentInformation, err := service.Paid(5000000)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n", paymentInformation)
}
