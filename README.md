# Protobuf vs. JSON+JSON Schema
This repository contains a comparison of Protobuf and JSON validated with
JSON-Schema.

It aims to quantify the following metrics:

1. Difference in 'protocol' overhead
1. Difference in per-metric overhead
1. Time to parse a small message
1. Time to parse a large message
1. Time to serialise a small message (including validation for JSON-Schema)
1. Time to serialise a large message (including validation for JSON-Schema)
1. If possible, how much memory is used for in-memory representation

The format used in this repository is not a final pipeline message format, but
rather a hypothetical one which is illustrative of format performance

## To run
Install dependencies

```bash
brew install go protobuf protoc-gen-go
```

```bash
go test -bench=. -v ./...
```

## Example
```
go test -bench=. -test.benchmem -test.benchtime=5s -v
=== RUN   TestSerialisedSizes
=== RUN   TestSerialisedSizes/PROTOBUF:_138_bytes
=== RUN   TestSerialisedSizes/JSON:_____378_bytes
=== RUN   TestSerialisedSizes/EMPTY_PROTOBUF:_6_bytes
=== RUN   TestSerialisedSizes/EMPTY_JSON:_____211_bytes
--- PASS: TestSerialisedSizes (0.00s)
    --- PASS: TestSerialisedSizes/PROTOBUF:_138_bytes (0.00s)
    --- PASS: TestSerialisedSizes/JSON:_____378_bytes (0.00s)
    --- PASS: TestSerialisedSizes/EMPTY_PROTOBUF:_6_bytes (0.00s)
    --- PASS: TestSerialisedSizes/EMPTY_JSON:_____211_bytes (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/Scout24/protobuf-vs-json
BenchmarkTester_SerialiseProtobuf-8   	20000000	       514 ns/op	5359273037.11 MB/s	     144 B/op	       1 allocs/op
BenchmarkTester_SerialiseJSON-8       	 5000000	      1834 ns/op	823149495.01 MB/s	     464 B/op	       4 allocs/op
BenchmarkTester_ParseProtobuf-8       	10000000	       783 ns/op	1761835970.23 MB/s	     504 B/op	      15 allocs/op
BenchmarkTester_ParseValidateJSON-8   	  200000	     34116 ns/op	2215920.31 MB/s	    6929 B/op	     222 allocs/op
BenchmarkTester_ParseJSON-8           	 1000000	      7081 ns/op	53380892.97 MB/s	     896 B/op	      21 allocs/op
PASS
ok  	github.com/Scout24/protobuf-vs-json	44.850s
```
