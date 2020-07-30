package order

import "coffee"

type CoffeeOrders []coffee.ICoffee

type Order struct {
	OrderId string
	Orders  []coffee.ICoffee
}

func NewOrder(orderId string, orders []coffee.ICoffee) *Order {
	return &Order{
		orderId, orders,
	}
}
