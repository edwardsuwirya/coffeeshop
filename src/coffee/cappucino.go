package coffee

import (
	"fmt"
	"sync"
	"time"
	"utils"
)

type cappucino struct {
	coffee       string
	milk         string
	orderStatus  *sync.WaitGroup
	updateStatus chan string
	orderId      string
	startTime    time.Time
	finishTime   time.Time
}

func (c *cappucino) Start(id string, status chan string, orderStatus *sync.WaitGroup) {
	//fmt.Printf("Start order %s : %v\n", id, time.Now())
	c.startTime = time.Now()
	info := fmt.Sprintf("Mulai buat cappucino, order :%s", id)
	c.updateStatus = status
	c.orderStatus = orderStatus
	c.orderId = id
	status <- info
	c.prepare()
}

func (c *cappucino) prepare() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		info := fmt.Sprintf(utils.FORMAT_GILING_KOPI, c.coffee)
		c.updateStatus <- info
		wg.Done()
	}()

	go func() {
		time.Sleep(2000 * time.Millisecond)
		info := fmt.Sprintf(utils.FORMAT_PANASKAN_SUSU, c.milk)
		c.updateStatus <- info
		wg.Done()
	}()

	wg.Wait()
	c.finish()
}

func (c *cappucino) finish() {
	time.Sleep(1000 * time.Millisecond)
	c.updateStatus <- utils.TUANG_CANGKIR
	c.finishTime = time.Now()
	info := fmt.Sprintf(utils.FORMAT_SAJIAN_CAPPUCINO, c.orderId, c.coffee, c.milk, c.startTime.Format(utils.FORMAT_TIME_STAMP), c.finishTime.Format(utils.FORMAT_TIME_STAMP))
	c.updateStatus <- info
	c.orderStatus.Done()
	//fmt.Printf("Finish order %s : %v\n", c.orderId, time.Now())

}

func NewCappucino(coffee string, milk string) ICoffee {
	return &cappucino{
		coffee,
		milk,
		nil, nil,
		"", time.Now(), time.Now(),
	}
}
