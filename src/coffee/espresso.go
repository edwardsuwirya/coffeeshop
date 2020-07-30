package coffee

import (
	"fmt"
	"orderStatus"
	"sync"
	"time"
	"utils"
)

type espresso struct {
	coffee             string
	orderProcessStatus orderStatus.OrderProsesStatus
}

func (c *espresso) Start(id string, status chan string, orderStatus *sync.WaitGroup) {
	//fmt.Printf("Start order %s : %v\n", id, time.Now())
	c.orderProcessStatus.StartTime = time.Now()
	info := fmt.Sprintf("Mulai buat Espresso, order :%s", id)
	c.orderProcessStatus.UpdateStatus = status
	c.orderProcessStatus.OrderStatus = orderStatus
	c.orderProcessStatus.OrderId = id
	status <- info
	c.prepare()
}

func (c *espresso) prepare() {
	time.Sleep(utils.GILING_ESPRESSO_TIME_LENGTH * time.Millisecond)
	info := fmt.Sprintf(utils.FORMAT_GILING_KOPI, c.coffee)
	c.orderProcessStatus.UpdateStatus <- info

	c.finish()
}

func (c *espresso) finish() {
	time.Sleep(utils.PENYAJIAN_CANGKIR_TIME_LENGTH * time.Millisecond)
	c.orderProcessStatus.UpdateStatus <- utils.TUANG_CANGKIR
	c.orderProcessStatus.FinishTime = time.Now()
	info := fmt.Sprintf(utils.FORMAT_SAJIAN_ESPRESSO, c.orderProcessStatus.OrderId, c.coffee, c.orderProcessStatus.StartTime.Format(utils.FORMAT_TIME_STAMP), c.orderProcessStatus.FinishTime.Format(utils.FORMAT_TIME_STAMP))
	c.orderProcessStatus.UpdateStatus <- info
	c.orderProcessStatus.OrderStatus.Done()
	//fmt.Printf("Finish order %s : %v\n", c.orderId, time.Now())

}

func NewEspresso(coffee string) ICoffee {
	return &espresso{
		coffee,
		orderStatus.NewOrderProcessStatus(),
	}
}
