package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
	"time"
)

var leakyStorage = make(map[int]*goquery.Document)
var mu = sync.Mutex{}

func generateHTML() string {
	var html string
	html = "<html><body>"
	for i := 0; i < 10000; i++ {
		html += fmt.Sprintf("<div>#{generateRandomString(10)}</div>")
	}
	html += "</body></html>"
	return html
}

func generateRandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	s := make([]rune, n)
	for i, _ := range s {
		s[i] = letters[rand.Int()%len(letters)]
	}
	return string(s)
}

func main() {
	cpuFile, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(cpuFile)

	defer cpuFile.Close()
	defer pprof.StopCPUProfile()

	for i := 0; i < 100; i++ {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(generateHTML()))
		if err != nil {
			fmt.Println("failed to parse HTML: ", err)
		}
		//simulate memory leak by storing the parsed documents
		mu.Lock()
		leakyStorage[time.Now().Nanosecond()] = doc
		mu.Unlock()
	}
	memoryFile, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(memoryFile)

	defer memoryFile.Close()

	fmt.Println("Profiling complete!")
}
