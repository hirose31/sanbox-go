package main

import (
	"fmt"
	"os"
	"runtime"
)

const name = "sandbox-go"
const version = "1.3.2"

var revision = "HEAD"

func hello(name string) (int, error) {
	fmt.Printf("Hello, %s!\n", name)

	if name == "hen" {
		return len(name), fmt.Errorf("%s", "okasii!")
	}

	return len(name), nil
}

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "--version" {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}
	len, err := hello("hogehoge")
	fmt.Printf("%d %#v\n", len, err)
}
