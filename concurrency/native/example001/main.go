package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Hello from goroutine")
}
func main() {

	go f()
	fmt.Println("Hello from main")
	time.Sleep(100)
}
