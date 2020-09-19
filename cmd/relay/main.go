package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jamesfcarter/pl2303relay"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: relay <device> <byte>\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}
	device := os.Args[1]
	value, err := strconv.Atoi(os.Args[2])
	if err != nil {
		usage()
	}
	relays, err := pl2303relay.New(device)
	if err != nil {
		panic(err)
	}
	if _, err := relays.Init(); err != nil {
		panic(err)
	}
	if err := relays.Update(byte(value)); err != nil {
		panic(err)
	}
}
