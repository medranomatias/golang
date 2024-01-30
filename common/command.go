package command

import (
	"flag"
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/pkg/profile"
)

type Empty struct {
}

func (empty Empty) Stop() {

}

func Parse() interface {
	Stop()
} {
	var deferCommand interface {
		Stop()
	}

	var cpuProfileRate = flag.Int("cpu-common-rate", math.MinInt, "")
	if *cpuProfileRate != math.MinInt {
		runtime.SetCPUProfileRate(*cpuProfileRate)
	}
	var mutexProfileFraction = flag.Int("mutex-common-fraction", math.MinInt, "")
	if *mutexProfileFraction != math.MinInt {
		runtime.SetMutexProfileFraction(*mutexProfileFraction)
	}
	var blockProfileRate = flag.Int("block-common-rate", math.MinInt, "")
	if *blockProfileRate != math.MinInt {
		runtime.SetBlockProfileRate(*blockProfileRate)
	}

	var memProfileRate = flag.Int("mem-common-rate", math.MinInt, "")
	if *memProfileRate != math.MinInt {
		runtime.MemProfileRate = *memProfileRate
	}
	var profileType string
	flag.StringVar(&profileType, "profile", "", "select common type")
	flag.Parse()
	switch strings.ToLower(profileType) {
	case "trace":

		deferCommand = profile.Start(profile.TraceProfile, profile.ProfilePath("."))
	case "block":
		deferCommand = profile.Start(profile.BlockProfile, profile.ProfilePath("."))
	case "mem":
		deferCommand = profile.Start(profile.MemProfile, profile.ProfilePath("."))
	case "alloc":
		deferCommand = profile.Start(profile.MemProfileAllocs, profile.ProfilePath("."))
	case "heap":
		deferCommand = profile.Start(profile.MemProfileHeap, profile.ProfilePath("."))
	case "mutex":
		deferCommand = profile.Start(profile.MutexProfile, profile.ProfilePath("."))
	case "clock":
		deferCommand = profile.Start(profile.ClockProfile, profile.ProfilePath("."))
	case "goroutine":
		deferCommand = profile.Start(profile.GoroutineProfile, profile.ProfilePath("."))
	case "cpu":
		deferCommand = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	default:
		fmt.Println("nothing to do")
		deferCommand = Empty{}
	}

	return deferCommand
}
