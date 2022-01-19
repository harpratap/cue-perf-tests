# cue-perf-tests

This repo contains some benchmarks that show parts of [CUE
](https://cuelang.org/) that ended up slowing down the export times of our cue config files

## 1. List Comprehensions

We make use of List comprehensions to convert from key-value maps to arrays.
The performance of export greatly depends on the number of such elements in the list and also on the size of the elements themselves.

Meaning iteration over 10 elements is costlier than iteration over 1 element
AND iteration of large definition like `{ a: int b: string }` is costlier over smaller definition like `string`

```
go test -bench=BenchmarkListComprehension -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/harpratap/cue-perf-tests
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkListComprehensionLong10-16     	    7579	    727905 ns/op
BenchmarkListComprehensionLong1-16      	   78135	     81278 ns/op
BenchmarkListComprehensionShort10-16    	   71486	     98007 ns/op
BenchmarkListComprehensionShort1-16     	  110869	     55206 ns/op
PASS
ok  	github.com/harpratap/cue-perf-tests	33.884s
```