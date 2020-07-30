package coffee

import (
	"fmt"
	"orderStatus"
	"sync"
	"time"
	"utils"
)

type cappucino struct {
	coffee             string
	milk               string
	orderProcessStatus orderStatus.OrderProsesStatus
}

func (c *cappucino) Start(id string, status chan string, orderStatus *sync.WaitGroup) {
	//fmt.Printf("Start order %s : %v\n", id, time.Now())
	c.orderProcessStatus.StartTime = time.Now()
	info := fmt.Sprintf("Mulai buat cappucino, order :%s", id)
	c.orderProcessStatus.UpdateStatus = status
	c.orderProcessStatus.OrderStatus = orderStatus
	c.orderProcessStatus.OrderId = id
	status <- info
	c.prepare()
}

func (c *cappucino) prepare() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(utils.GILING_KOPI_TIME_LENGTH * time.Millisecond)
		info := fmt.Sprintf(utils.FORMAT_GILING_KOPI, c.coffee)
		c.orderProcessStatus.UpdateStatus <- info
		wg.Done()
	}()

	go func() {
		time.Sleep(utils.PANASKAN_SUSU_TIME_LENGTH * time.Millisecond)
		info := fmt.Sprintf(utils.FORMAT_PANASKAN_SUSU, c.milk)
		c.orderProcessStatus.UpdateStatus <- info
		wg.Done()
	}()

	wg.Wait()
	c.finish()
}

func (c *cappucino) finish() {
	time.Sleep(utils.PENYAJIAN_CANGKIR_TIME_LENGTH * time.Millisecond)
	c.orderProcessStatus.UpdateStatus <- utils.TUANG_CANGKIR
	c.orderProcessStatus.FinishTime = time.Now()
	info := fmt.Sprintf(utils.FORMAT_SAJIAN_CAPPUCINO, c.orderProcessStatus.OrderId, c.coffee, c.milk, c.orderProcessStatus.StartTime.Format(utils.FORMAT_TIME_STAMP), c.orderProcessStatus.FinishTime.Format(utils.FORMAT_TIME_STAMP))
	c.orderProcessStatus.UpdateStatus <- info
	c.orderProcessStatus.OrderStatus.Done()
	//fmt.Printf("Finish order %s : %v\n", c.orderId, time.Now())

}

func NewCappucino(coffee string, milk string) ICoffee {
	return &cappucino{
		coffee,
		milk,
		orderStatus.NewOrderProcessStatus(),
	}
}
