# Go Timestamp Counter Benchmarks

Benchmarks to determine fastest way to have a shared gauge containing a timestamp; where writes
are frequent but reads are somewhat infrequent.

Structure largely copied from: https://gist.github.com/quasilyte/009edaf14aad08f6d1997b026c63c0a0

```shell
❯ date
Fri Oct  2 13:56:08 CDT 2020

go-ts-counter-benchmarks on  master [?] via 🐹 v1.15.2
❯ go test --bench=. --benchtime=20s --count=4
goos: darwin
goarch: amd64
pkg: github.com/dudleycodes/go-ts-counter-benchmarks
BenchmarkAtomic-12      71911761               366 ns/op
BenchmarkAtomic-12      59503452               404 ns/op
BenchmarkAtomic-12      57908635               410 ns/op
BenchmarkAtomic-12      60280071               384 ns/op
BenchmarkMutex-12       52983196               445 ns/op
BenchmarkMutex-12       53645341               452 ns/op
BenchmarkMutex-12       52201502               454 ns/op
BenchmarkMutex-12       53146382               452 ns/op
BenchmarkRWMutex-12     40341808               594 ns/op
BenchmarkRWMutex-12     40784523               595 ns/op
BenchmarkRWMutex-12     40465648               594 ns/op
BenchmarkRWMutex-12     40271937               592 ns/op
```
