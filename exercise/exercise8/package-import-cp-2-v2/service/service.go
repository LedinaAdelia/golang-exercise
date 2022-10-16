package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	a, err := s.database.GetProductByname(productName)
	if err != nil {
		return err
	}
	if quantity < 1 {
		return errors.New("invalid quantity")
	}
	cart, err := s.database.Load()
	if err != nil {
		return err
	}
	cart = append(cart, entity.CartItem{
		ProductName: a.Name,
		Price:       a.Price,
		Quantity:    quantity,
	})
	return s.database.Save(cart)
}

func (s *Service) RemoveCart(productName string) error {
	a, err := s.database.GetProductByname(productName)
	try := false
	if err != nil {
		return err
	}
	cart, err := s.database.Load()
	if err != nil {
		return err
	}
	for i := 0; i < len(cart); i++ {
		if cart[i].ProductName == a.Name {
			cart[i] = cart[len(cart)-1]
			try = true
			return s.database.Save(cart[:len(cart)-1])
		}
	}
	if try == false {
		return errors.New("product not found")
	}
	return nil
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (s *Service) ResetCart() error {
	cart, err := s.database.Load()
	if err != nil {
		return err
	}
	cart = nil
	return s.database.Save(cart) // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	list := database.NewDatabase().GetProductData()
	return list, nil // TODO: replace this
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {
	cart, err := s.database.Load()
	total := 0
	if err != nil {
		return entity.PaymentInformation{}, err
	}
	for _, v := range cart {
		total += v.Price * v.Quantity
	}
	change := money - total
	zonk := entity.PaymentInformation{}
	if change < 0 {
		return zonk, errors.New("money is not enough")
	}
	try := entity.PaymentInformation{
		ListProduct: cart,
		TotalPrice:  total,
		MoneyPaid:   money,
		Change:      change,
	}
	return try, s.ResetCart() // TODO: replace this
}
