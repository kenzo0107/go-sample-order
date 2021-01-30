package main

import "github.com/kenzo0107/go-sample-order/pkg"

var sampleVar = defaultVar()

func defaultVar() int {
	println(pkg.SampleVar)
	println("main.var")
	return 1
}

func init() {
	println("main.init")
}

func main() {
	println("main.main")
}
