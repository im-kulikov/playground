package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"unicode"

	"github.com/pkg/profile"
	"golang.org/x/exp/mmap"
)

var (
	_ = readByte
	_ = profile.Profile{}
)

func readByte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	start := time.Now()
	defer func() {
		log.Print(time.Since(start))
	}()
	defer profile.Start().Stop()
	f, err := mmap.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	var (
		array  [1]byte
		offset int64
		length = int64(f.Len())
		words  = 0
		inword = false
	)

	for offset = 0; offset < length; offset++ {
		_, err := f.ReadAt(array[:], offset)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}

		r := rune(array[0])
		tmp := unicode.IsLetter(r)
		if !tmp && inword {
			words++
		}
		inword = tmp
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
