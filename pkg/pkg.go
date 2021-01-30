package pkg

// SampleVar : sample variable in pkg
var SampleVar = defaultVar()

func defaultVar() int {
	println("pkg.var")
	return 1
}

func init() {
	println("pkg.init")
}
