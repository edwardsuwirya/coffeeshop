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
		if strings.HasPrefix(s, utils.SAJIAN) {
			var s = strings.Split(s, ",")
			var o = strings.Split(s[0], utils.SAJIAN+" ")
			fmt.Printf(utils.FORMAT_STRING_HEADER, utils.Trimming(o[1]), utils.Trimming(s[1]), utils.Trimming(s[2])+" "+utils.Trimming(s[3]))
		}
	}
	utils.PrintBorder()
}

func NewKitchenProduction(orders *[]order.Order) *production {
	return &production{orders}
}
