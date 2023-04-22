#### Benchmarks

RLP (go-ethereum structs)

```go
$ go test . -bench .                                                                                                                       ✔
goos: darwin
goarch: arm64
pkg: scratch/fastrlp

BenchmarkRLPHeader-8                     3808970               312.1 ns/op
BenchmarkFastRLPHeader-8                 2504481               460.9 ns/op
BenchmarkMsgpackHeader-8                 1813960               662.8 ns/op
BenchmarkJSONHeader-8                     472345              2431.0 ns/op

BenchmarkRLPReceipt-8                    2431597               494.4 ns/op
BenchmarkFastRLPReceipt-8                3547752               337.9 ns/op

BenchmarkRLPTransaction-8                6527433               181.8 ns/op
BenchmarkFastRLPTransaction-8            4475868               264.9 ns/op

PASS
ok      scratch/fastrlp 12.460s
```

SSZ (custom struct [**not** go-ethereum])

```go
$ go test . -bench .
goos: darwin
goarch: arm64
pkg: scratch/fastssz

BenchmarkRLPHeader-8             3831055               306.0 ns/op
BenchmarkFastRLPHeader-8         2689016               433.9 ns/op
BenchmarkSSZHeader-8            12150908                95.0 ns/op

PASS
ok      scratch/fastssz 4.964s
```
