package coffee

import (
	"fmt"
	"sync"
	"time"
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
		info := fmt.Sprintf("Giling kopi :%s", c.coffee)
		c.updateStatus <- info
		wg.Done()
	}()

	go func() {
		time.Sleep(2000 * time.Millisecond)
		info := fmt.Sprintf("Panaskan susu :%s", c.milk)
		c.updateStatus <- info
		wg.Done()
	}()

	wg.Wait()
	c.finish()
}

func (c *cappucino) finish() {
	time.Sleep(1000 * time.Millisecond)
	c.updateStatus <- "Tuangkan ke cangkir"
	c.finishTime = time.Now()
	info := fmt.Sprintf("Menyajikan %s, Kopi Capucino %s %s, mulai %v, selesai %v", c.orderId, c.coffee,c.milk, c.startTime.Format("2006-01-02 15:04:05"), c.finishTime.Format("2006-01-02 15:04:05"))
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
