package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/pkg/profile"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}
func main() {
	ExampleStart_withFlags()
	for _, s := range []string{"a", "b", "c"} {
		go f(s)
	}
	time.Sleep(100)
}

func ExampleStart_withFlags() {
	// use the flags package to selectively enable profiling.
	mode := flag.String("tools.mode", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
	flag.Parse()
	fmt.Printf("Goroutine %s\n", *mode)
	switch *mode {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "mutex":
		defer profile.Start(profile.MutexProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	default:
		// do nothing
	}
}
