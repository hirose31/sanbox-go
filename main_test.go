package main

import "testing"

func TestHello(t *testing.T) {
	_, err := hello("hogehoge")
	if err != nil {
		t.Errorf("hogehoge returns error")
	}

	_, err = hello("hen")
	if err == nil {
		t.Errorf("hen does not return error")
	}
}
