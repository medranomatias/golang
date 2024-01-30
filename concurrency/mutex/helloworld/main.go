package main

import (
	"fmt"
	"github.com/medranomatias/common"
	"sync"
)

func main() {
	defer command.Parse().Stop()
	var m, p, q sync.Mutex

	p.Lock()
	q.Lock()

	go func() {
		m.Lock()
		fmt.Println("critial section f")
		m.Unlock()
		q.Unlock()
	}()

	go func() {
		m.Lock()
		fmt.Println("critial section g")
		m.Unlock()
		p.Unlock()
	}()

	p.Lock()
	q.Lock()
	fmt.Println("release lock")
}
