package kitchen

import (
	"fmt"
	"order"
	"strings"
	"sync"
	"utils"
)

type production struct {
	orders *[]order.Order
}

func (p *production) Produce() {
	status := make(chan string)

	var orderStatus sync.WaitGroup
	for _, o := range *(p.orders) {
		for _, oc := range o.Orders {
			orderStatus.Add(1)
			go oc.Start(o.OrderId, status, &orderStatus)
		}
	}

	//IIFE
	go func() {
		orderStatus.Wait()
		close(status)
	}()

	//<-status
	utils.PrintHeader()
	for s := range status {
		if strings.HasPrefix(s, "Menyajikan") {
			var s = strings.Split(s, ",")
			var o = strings.Split(s[0], "Menyajikan ")
			fmt.Printf("%-10s %-50s %-20s\n", utils.Trimming(o[1]), utils.Trimming(s[1]), utils.Trimming(s[2])+" "+utils.Trimming(s[3]))
		}
	}
	utils.PrintBorder()
}

func NewCoffeeMaker(orders *[]order.Order) *production {
	return &production{orders}
}
