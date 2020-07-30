package orderStatus

import (
	"sync"
	"time"
)

type OrderProsesStatus struct {
	OrderStatus  *sync.WaitGroup
	UpdateStatus chan string
	OrderId      string
	StartTime    time.Time
	FinishTime   time.Time
}

func NewOrderProcessStatus() OrderProsesStatus {
	return OrderProsesStatus{
		OrderStatus:  nil,
		UpdateStatus: nil,
		OrderId:      "",
		StartTime:    time.Now(),
		FinishTime:   time.Now(),
	}
}
