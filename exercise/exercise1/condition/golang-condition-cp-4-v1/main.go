package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	totalTicket := VIP + regular + student
	priceVIP := 30 * VIP
	priceRegular := 20 * regular
	priceStudent := 10 * student
	totalPrice := priceVIP + priceRegular + priceStudent

	if day%2 == 0 && totalPrice >= 100 && totalTicket < 5 {
		coba := float32(totalPrice) * 0.10
		return float32(totalPrice) - coba
	} else if day%2 == 0 && totalPrice >= 100 && totalTicket >= 5 {
		coba := float32(totalPrice) * 0.20
		return float32(totalPrice) - coba
	} else if totalPrice >= 100 && totalTicket < 5 {
		coba := float32(totalPrice) * 0.15
		return float32(totalPrice) - coba
	} else if totalPrice >= 100 && totalTicket >= 5 {
		coba := float32(totalPrice) * 0.25
		return float32(totalPrice) - coba
	} else {
		return float32(totalPrice)
	}

	return 0.0 // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
