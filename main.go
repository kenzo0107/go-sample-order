package main

import "github.com/kenzo0107/go-sample-order/pkg"

// samepleVar : sample variable
var sampleVar = defaultVar()

func defaultVar() string {
	println(pkg.SampleVar)
	println("main.var")
	return ""
}

func init() {
	println(sampleVar)
	println("main.init")
}

func main() {
	println("main.main")
}
