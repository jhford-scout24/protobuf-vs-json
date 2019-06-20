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
go test -bench=.
```

EXAMPLE:
```
=== RUN   TestSerialisedSizes
=== RUN   TestSerialisedSizes/PROTOBUF:_138_bytes
=== RUN   TestSerialisedSizes/JSON:_____378_bytes
--- PASS: TestSerialisedSizes (0.00s)
    --- PASS: TestSerialisedSizes/PROTOBUF:_138_bytes (0.00s)
    --- PASS: TestSerialisedSizes/JSON:_____378_bytes (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/Scout24/protobuf-vs-json
BenchmarkTester_SerialiseProtobuf-8   	 3000000	       518 ns/op	 266.35 MB/s
BenchmarkTester_SerialiseJSON-8       	 1000000	      1854 ns/op	 162.87 MB/s
BenchmarkTester_ParseProtobuf-8       	 2000000	       780 ns/op	 176.82 MB/s
BenchmarkTester_ParseValidateJSON-8   	   50000	     33001 ns/op	  11.45 MB/s
BenchmarkTester_ParseJSON-8           	  200000	      7031 ns/op	  53.76 MB/s
BenchmarkTester_ValidateJSON-8        	   50000	     25360 ns/op	  14.90 MB/s
PASS
ok  	github.com/Scout24/protobuf-vs-json	11.331s
```
