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
		*(NewOrder("123", CoffeeOrders{order1, order3})),
		*(NewOrder("234", CoffeeOrders{order2})),
	}
	return &orders
}
