package main

import (
	"coffee"
	"fmt"
	"strings"
	"sync"
)

type order struct {
	orderId string
	orders  []coffee.ICoffee
}

func main() {
	var order1 = coffee.NewCappucino("Kopi Lampung", "Low Fat")
	var order2 = coffee.NewCappucino("Kopi Luwak", "Full Cream")
	var order3 = coffee.NewEspresso("Kopi sumatra")
	var orders = []order{
		{
			orderId: "123",
			orders: []coffee.ICoffee{
				order1, order3,
			},
		},
		{
			orderId: "234",
			orders: []coffee.ICoffee{
				order2,
			},
		},
	}
	status := make(chan string)

	var orderStatus sync.WaitGroup
	for _, o := range orders {
		for _, oc := range o.orders {
			orderStatus.Add(1)
			go oc.Start(o.orderId, status, &orderStatus)
		}
	}

	//IIFE
	go func() {
		orderStatus.Wait()
		close(status)
	}()

	//<-status
	printHeader()
	for s := range status {
		if (strings.HasPrefix(s, "Menyajikan")) {
			var s = strings.Split(s, ",")
			var o = strings.Split(s[0], "Menyajikan ")
			fmt.Printf("%-10s %-50s %-20s\n", o[1], s[1], s[2]+" "+s[3])
		}
	}
	printBorder()

}

func printHeader() {
	printBorder()
	fmt.Printf("%-10s %-50s %-20s\n", "Order", "Item", "Serving Time")
	printBorder()

}
func printBorder() {
	fmt.Printf("%120s\n", strings.Repeat("=", 120))
}
