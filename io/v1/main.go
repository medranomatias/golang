package main

import (
	"fmt"
	"github.com/pkg/profile"
	"io"
	"log"
	"os"
	"unicode"
)

func readbyte(r io.Reader) (rune, error) {
	var buff [1]byte
	_, err := r.Read(buff[:])
	return rune(buff[0]), err
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("couldnt open file %q: %v", os.Args[1], err)
	}
	words := 0
	inword := false
	for {
		r, err2 := readbyte(f)
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatalf("couldn't read the file %q %v", os.Args[1], err2)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words", os.Args[1], words)
}
