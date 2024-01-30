package main

import (
	"fmt"
	command "github.com/medranomatias/common"
	"time"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}
func main() {
	defer command.Parse().Stop()
	for _, s := range []string{"a", "b", "c"} {
		go func() {
			fmt.Printf("Goroutine %s\n", s)
		}()
	}
	time.Sleep(100)
}
