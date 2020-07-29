package coffee

import "sync"

type ICoffee interface {
	Start(id string, status chan string, orderStatus *sync.WaitGroup)
	prepare()
	finish()
}
