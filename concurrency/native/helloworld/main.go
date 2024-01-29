package main

import (
	"fmt"
	"github.com/pkg/profile"
	"time"
)

func f() {
	fmt.Println("Hello from goroutine")
}
func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	go f()
	fmt.Println("Hello from main")
	time.Sleep(100)
}
