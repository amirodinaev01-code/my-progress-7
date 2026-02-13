package main

import (
	"fmt"
)

type Payment interface {
	Wallet()
}

type Order struct {
	Name  string
	Price float64
	Pay   float64
}

func (p Order) Wallet() {
	fmt.Printf("Payment successful for %s. Enjoy!\n", p.Name)
}

type Validator func(Order) bool

func main() {
	order := Order{Name: "Barbeq", Price: 120.0, Pay: 500.0}

	examination := []Validator{
		func(o Order) bool { return o.Pay >= o.Price },
	}

	if LogikDish(order, examination) {
		PayService(order)
	}
}

func PayService(p Payment) {
	p.Wallet()
}

func LogikDish(o Order, check []Validator) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: recovered from panic")
			ok = false
		} else if !ok {
			fmt.Println("Transaction declined")
		} else {
			fmt.Println("Validation passed")
		}
	}()

	for _, checkFunc := range check {
		if !checkFunc(o) {
			return false
		}
	}

	ok = true
	return ok
}