package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	name = kingpin.Flag("name", "Name of you.").Default("NONAME").String()
)

func main() {
	kingpin.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}
