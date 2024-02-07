package main

import (
	"github.com/medranomatias/common"
)

func main() {
	defer command.Parse().Stop()
	b := make([]byte, 600*1024*1024)
	b[5000] = 1
	b[100000] = 1
	b[104000] = 1
	for i := range b {
		b[i] = 1
	}
}
