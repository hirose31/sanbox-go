package main

import (
	"fmt"
)

func hello(name string) (int, error) {
	fmt.Printf("Hello, %s!\n", name)

	if name == "hen" {
		return len(name), fmt.Errorf("%s", "okasii!")
	}

	return len(name), nil
}

func main() {
	len, err := hello("hogehoge")
	fmt.Printf("%d %#v\n", len, err)
}
