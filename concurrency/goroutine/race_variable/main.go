// go tool compile "-m" main.go #check if variable move to heap or still on stack
package main

import (
	"fmt"
	command "github.com/medranomatias/common"
	"time"
)

func main() {
	defer command.Parse().Stop()
	for _, s := range []string{"a", "b", "c"} {
		go func() {
			fmt.Printf("Goroutine %s\n", s)
		}()
	}
	time.Sleep(100)
}
