# Study of `mat` import effect on code size and compilation speed

## No `mat` import
Vendor lines of code: 0

Binary size is 1.3M.

#### Go mod tidy
```
$ go clean -cache -modcache
$ time go mod tidy

real    0m0,016s
user    0m0,011s
sys     0m0,004s
```

#### Go build
```
$ go clean -cache
$ time go build .

real    0m0,140s
user    0m0,131s
sys     0m0,046s
```

## With `mat` import
Vendor lines of code: 52476
Of which 46575 are Go code.

Binary size is 1.5M.

#### Go mod tidy
```
$ go clean -cache -modcache
$ time go mod tidy
go: finding module for package gonum.org/v1/gonum/mat
go: downloading gonum.org/v1/gonum v0.11.0
go: found gonum.org/v1/gonum/mat in gonum.org/v1/gonum v0.11.0
go: downloading golang.org/x/exp v0.0.0-20191002040644-a1355ae1e2c3

real    0m11,027s
user    0m0,685s
sys     0m0,299s
```

#### Go build

```
$ go clean -cache
$ time go build -tags=matty .

real    0m2,436s
user    0m7,189s
sys     0m0,486s
```