package main

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}
func main() {
	defer profileOption().Stop()
	for _, s := range []string{"a", "b", "c"} {
		go f(s)
	}
	time.Sleep(100)
}
