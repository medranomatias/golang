package main

import (
	"fmt"
	"github.com/medranomatias/common"
	"time"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}
func main() {
	defer command.Parse()
	for _, s := range []string{"a", "b", "c"} {
		go f(s)
	}
	time.Sleep(100)
}
