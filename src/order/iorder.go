package order

import "coffee"

type IOrder interface {
	GetOrders() *[]Order
}

type DummyOrder struct {
}

func (do *DummyOrder) GetOrders() *[]Order {
	var order1 = coffee.NewCappucino("Kopi Lampung", "Low Fat")
	var order2 = coffee.NewCappucino("Kopi Luwak", "Full Cream")
	var order3 = coffee.NewEspresso("Kopi sumatra")
	var orders = []Order{
		{
			OrderId: "123",
			Orders: []coffee.ICoffee{
				order1, order3,
			},
		},
		{
			OrderId: "234",
			Orders: []coffee.ICoffee{
				order2,
			},
		},
	}
	return &orders
}
