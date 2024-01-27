package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var data = make(map[int]int)
var mu sync.Mutex
var wg sync.WaitGroup

func slowFunction() {
	time.Sleep(500 * time.Millisecond)
}

func writeData() {
	mu.Lock()
	defer mu.Unlock()
	data[1] = 1
	slowFunction()
}

func readData() int {
	mu.Lock()
	defer mu.Unlock()
	return data[1]
}

func main() {
	file, err := os.Create("mutex.prof")
	runtime.SetMutexProfileFraction(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		go writeData()
	}
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 10; i++ {
		go fmt.Println(readData())
	}
	pprof.Lookup("mutex").WriteTo(file, 0)
	wg.Wait()
	defer file.Close()
	var v io.Writer
}
