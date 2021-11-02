package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"os"
	"strings"
)

func main() {
	var s string
	flag.StringVar(&s, "s", "", "input string")
	flag.Parse()

	count := 0
	var r io.Reader
	if s != "" {
		r = strings.NewReader(s)
	} else {
		count -= 1
		r = bufio.NewReader(os.Stdin)
	}

	buf := make([]byte, 1<<20)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		count += len([]rune(string(buf[:n])))
	}

	fmt.Printf("%s\n", humanize.Comma(int64(count)))
}
