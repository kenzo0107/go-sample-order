This repository is created to check the processing order in Go.

Result:

```
$ go test

pkg.var
pkg.init
1
main.var
2
main.init
test.init
test.setup
testing: warning: no tests to run
PASS
ok      github.com/kenzo0107/go-sample-order    0.272s [no tests to run]
?       github.com/kenzo0107/go-sample-order/pkg        [no test files]
```
