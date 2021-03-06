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

## 2. Conjunctions

Conjunctions are like AND operators used for specifying multiple constraints. But using too many constraints slows down the export timings too. Meaning `#level1 & { text1: text2: "hello" }` is less costly than `#level1 & { text1: #level2 & { text2: "hello" }` even though both behave exactly the same.

```
go test -bench=BenchmarkConjunction -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/harpratap/cue-perf-tests
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkConjunctionDouble-16    	   17859	    335044 ns/op
BenchmarkConjunctionSingle-16    	   19837	    287302 ns/op
PASS
ok  	github.com/harpratap/cue-perf-tests	25.784s
```

## 3. Openness

Definitions are by default closed, but adding ellipses `...` will open them. When importing OpenAPI spec to CUE it defaults to open definitions, which will cause export times to slow down

```
go test -bench=BenchmarkDefinition
goos: darwin
goarch: amd64
pkg: github.com/harpratap/cue-perf-tests
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkDefinitionOpen-16      	    1404	    800391 ns/op
BenchmarkDefinitionClosed-16    	    1635	    770721 ns/op
PASS
ok  	github.com/harpratap/cue-perf-tests	3.130s
```

## 4. Defaulting

Defaulting allows to provide a value to a type which can be overridden but only once. Adding more defaults also cause slowness

```
go test -bench=BenchmarkDefault -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/harpratap/cue-perf-tests
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkDefault0-16     	   41304	    144890 ns/op
BenchmarkDefault10-16    	   30867	    197449 ns/op
PASS
ok  	github.com/harpratap/cue-perf-tests	20.813s
```