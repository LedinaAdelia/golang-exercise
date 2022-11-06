package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (api *API) AddCart(w http.ResponseWriter, r *http.Request) {
	log.Println("hehe")
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error: "Request Product Not Found",
		})
		return
	}

	user, _ := api.sessionsRepo.ReadSessions()
	c, _ := r.Cookie("session_token")
	username := ""
	for _, v := range user {
		if v.Token == c.Value {
			username = v.Username
		}
	}
	r.ParseForm()
	var list []model.Product
	for _, formList := range r.Form {
		for _, v := range formList {
			item := strings.Split(v, ",")
			p, _ := strconv.ParseFloat(item[2], 64)
			q, _ := strconv.ParseFloat(item[3], 64)
			total := p * q
			list = append(list, model.Product{
				Id:       item[0],
				Name:     item[1],
				Price:    item[2],
				Quantity: item[3],
				Total:    total,
			})
		}
	}
	total := 0.0
	for _, v := range list {
		total += v.Total
	}
	// Add data field Name, Cart and TotalPrice with struct model.Cart.
	cart := model.Cart{
		Name:       username,
		Cart:       list,
		TotalPrice: total,
	} // TODO: replace this

	_, err := api.cartsRepo.CartUserExist(cart.Name)
	if err != nil {
		api.cartsRepo.AddCart(cart)
	} else {
		api.cartsRepo.UpdateCart(cart)
	}
	api.dashboardView(w, r)
}
