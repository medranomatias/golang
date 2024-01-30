package main

import (
	"fmt"
	"github.com/medranomatias/common"
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
