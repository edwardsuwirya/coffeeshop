package coffee

import (
	"fmt"
	"sync"
	"time"
	"utils"
)

type espresso struct {
	coffee       string
	orderStatus  *sync.WaitGroup
	updateStatus chan string
	orderId      string
	startTime    time.Time
	finishTime   time.Time
}

func (c *espresso) Start(id string, status chan string, orderStatus *sync.WaitGroup) {
	//fmt.Printf("Start order %s : %v\n", id, time.Now())
	c.startTime = time.Now()
	info := fmt.Sprintf("Mulai buat Espresso, order :%s", id)
	c.updateStatus = status
	c.orderStatus = orderStatus
	c.orderId = id
	status <- info
	c.prepare()
}

func (c *espresso) prepare() {

	time.Sleep(4000 * time.Millisecond)
	info := fmt.Sprintf(utils.FORMAT_GILING_KOPI, c.coffee)
	c.updateStatus <- info

	c.finish()
}

func (c *espresso) finish() {
	time.Sleep(1000 * time.Millisecond)
	c.updateStatus <- utils.TUANG_CANGKIR

	c.finishTime = time.Now()
	info := fmt.Sprintf(utils.FORMAT_SAJIAN_ESPRESSO, c.orderId, c.coffee, c.startTime.Format(utils.FORMAT_TIME_STAMP), c.finishTime.Format(utils.FORMAT_TIME_STAMP))
	c.updateStatus <- info
	c.orderStatus.Done()
	//fmt.Printf("Finish order %s : %v\n", c.orderId, time.Now())

}

func NewEspresso(coffee string) ICoffee {
	return &espresso{
		coffee,
		nil, nil,
		"", time.Now(), time.Now(),
	}
}
