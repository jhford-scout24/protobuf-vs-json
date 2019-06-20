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
BenchmarkTester_SerialiseProtobuf-8   	 3000000	       511 ns/op	 269.78 MB/s
BenchmarkTester_SerialiseJSON-8       	 1000000	      1813 ns/op	 166.56 MB/s
BenchmarkTester_ParseProtobuf-8       	 2000000	       768 ns/op	 179.57 MB/s
BenchmarkTester_ParseValidateJSON-8   	   50000	     32933 ns/op	  11.48 MB/s
BenchmarkTester_ParseJSON-8           	  200000	      6891 ns/op	  54.85 MB/s
BenchmarkTester_ValidateJSON-8        	   50000	     24668 ns/op	  15.32 MB/s
PASS
ok  	github.com/Scout24/protobuf-vs-json	11.139s
```
