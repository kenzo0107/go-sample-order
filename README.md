This repository is created to check the processing order in Go.

Result:

```
$ go test

pkg.var
pkg.init
1
main.var
main.init
test.init
test.setup
testing: warning: no tests to run
PASS
ok      github.com/kenzo0107/go-sample-order    0.271s
```
