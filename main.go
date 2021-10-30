package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	fmt.Printf("%d\n", len([]rune(s)))
}
