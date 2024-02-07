package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Hello from goroutine")
}
func main() {
	defer command.Parse()
	go f()
	fmt.Println("Hello from main")
	time.Sleep(100)
}
