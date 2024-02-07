package main

import (
	"fmt"
	"time"
)

func main() {
	defer command.Parse().Stop()
	var chn chan bool = make(chan bool)
	go func() {
		chn <- true
	}()
	go func() {
		var y bool = <-chn
		fmt.Println(y)
	}()
	time.Sleep(100)
}
