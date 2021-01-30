package main

import "testing"

func init() {
	println("test.init")
}

func setup() {
	println("test.setup")
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}
