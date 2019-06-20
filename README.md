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
go test -bench=. -test.benchmem
goos: darwin
goarch: amd64
pkg: github.com/Scout24/protobuf-vs-json
BenchmarkTester_SerialiseProtobuf-8   	 3000000	       513 ns/op	806593271.99 MB/s	     144 B/op	       1 allocs/op
BenchmarkTester_SerialiseJSON-8       	 1000000	      1839 ns/op	164147781.65 MB/s	     464 B/op	       4 allocs/op
BenchmarkTester_ParseProtobuf-8       	 2000000	       772 ns/op	357396783.61 MB/s	     504 B/op	      15 allocs/op
BenchmarkTester_ParseValidateJSON-8   	   50000	     32521 ns/op	581144.78 MB/s	    6753 B/op	     216 allocs/op
BenchmarkTester_ParseJSON-8           	  200000	      6894 ns/op	10964482.48 MB/s	     895 B/op	      21 allocs/op
BenchmarkTester_ValidateJSON-8        	   50000	     24572 ns/op	769133.76 MB/s	    5841 B/op	     194 allocs/op
PASS
ok  	github.com/Scout24/protobuf-vs-json	11.157s
```
