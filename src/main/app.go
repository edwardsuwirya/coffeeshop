package main

import (
	"kitchen"
	"order"
)

type coffeeShopApp struct {
	orders *[]order.Order
}

func newApp(orders order.IOrder) *coffeeShopApp {
	return &coffeeShopApp{
		orders: orders.GetOrders(),
	}
}
func (csp coffeeShopApp) run() {
	kitchen.NewCoffeeMaker(csp.orders).Produce()
}
func main() {
	newApp(&order.DummyOrder{}).run()
}
