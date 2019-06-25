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
go test -bench=. -test.benchmem -v -timeout 20m
=== RUN   TestSerialisedSizes
=== RUN   TestSerialisedSizes/fields,0,protobuf,size,81
=== RUN   TestSerialisedSizes/fields,0,json,size,248
=== RUN   TestSerialisedSizes/metrics,15,protobuf,size,786
=== RUN   TestSerialisedSizes/metrics,15,json,size,1546
=== RUN   TestSerialisedSizes/fields,15,protobuf,size,2601
=== RUN   TestSerialisedSizes/fields,15,json,size,3389
=== RUN   TestSerialisedSizes/metrics,0,protobuf,size,81
=== RUN   TestSerialisedSizes/metrics,0,json,size,248
=== RUN   TestSerialisedSizes/2kb_msg,protobuf,size,2132
=== RUN   TestSerialisedSizes/2kb_msg,json,size,2684
=== RUN   TestSerialisedSizes/fields,1,protobuf,size,249
=== RUN   TestSerialisedSizes/fields,1,json,size,467
=== RUN   TestSerialisedSizes/fields,3,protobuf,size,585
=== RUN   TestSerialisedSizes/fields,3,json,size,912
=== RUN   TestSerialisedSizes/fields,127,protobuf,size,21417
=== RUN   TestSerialisedSizes/fields,127,json,size,26916
=== RUN   TestSerialisedSizes/fields,1023,protobuf,size,171945
=== RUN   TestSerialisedSizes/fields,1023,json,size,216650
=== RUN   TestSerialisedSizes/small_msg,protobuf,size,106
=== RUN   TestSerialisedSizes/small_msg,json,size,271
=== RUN   TestSerialisedSizes/fields,63,protobuf,size,10665
=== RUN   TestSerialisedSizes/fields,63,json,size,13704
=== RUN   TestSerialisedSizes/metrics,127,protobuf,size,6050
=== RUN   TestSerialisedSizes/metrics,127,json,size,11231
=== RUN   TestSerialisedSizes/16kb_msg,protobuf,size,16469
=== RUN   TestSerialisedSizes/16kb_msg,json,size,19729
=== RUN   TestSerialisedSizes/fields,31,protobuf,size,5289
=== RUN   TestSerialisedSizes/fields,31,json,size,6730
=== RUN   TestSerialisedSizes/metrics,7,protobuf,size,410
=== RUN   TestSerialisedSizes/metrics,7,json,size,852
=== RUN   TestSerialisedSizes/metrics,63,protobuf,size,3042
=== RUN   TestSerialisedSizes/metrics,63,json,size,5784
=== RUN   TestSerialisedSizes/no_msg,protobuf,size,81
=== RUN   TestSerialisedSizes/no_msg,json,size,248
=== RUN   TestSerialisedSizes/fields,7,protobuf,size,1257
=== RUN   TestSerialisedSizes/fields,7,json,size,1724
=== RUN   TestSerialisedSizes/metrics,3,protobuf,size,222
=== RUN   TestSerialisedSizes/metrics,3,json,size,513
=== RUN   TestSerialisedSizes/metrics,31,protobuf,size,1538
=== RUN   TestSerialisedSizes/metrics,31,json,size,2898
=== RUN   TestSerialisedSizes/metrics,255,protobuf,size,12066
=== RUN   TestSerialisedSizes/metrics,255,json,size,22317
=== RUN   TestSerialisedSizes/8kb_msg,protobuf,size,8276
=== RUN   TestSerialisedSizes/8kb_msg,json,size,9921
=== RUN   TestSerialisedSizes/fields,255,protobuf,size,42921
=== RUN   TestSerialisedSizes/fields,255,json,size,54159
=== RUN   TestSerialisedSizes/fields,511,protobuf,size,85929
=== RUN   TestSerialisedSizes/fields,511,json,size,108538
=== RUN   TestSerialisedSizes/metrics,1,protobuf,size,128
=== RUN   TestSerialisedSizes/metrics,1,json,size,344
=== RUN   TestSerialisedSizes/metrics,511,protobuf,size,24098
=== RUN   TestSerialisedSizes/metrics,511,json,size,44621
=== RUN   TestSerialisedSizes/metrics,1023,protobuf,size,48162
=== RUN   TestSerialisedSizes/metrics,1023,json,size,89023
=== RUN   TestSerialisedSizes/4kb_msg,protobuf,size,4180
=== RUN   TestSerialisedSizes/4kb_msg,json,size,5154
--- PASS: TestSerialisedSizes (0.01s)
    --- PASS: TestSerialisedSizes/fields,0,protobuf,size,81 (0.00s)
    --- PASS: TestSerialisedSizes/fields,0,json,size,248 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,15,protobuf,size,786 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,15,json,size,1546 (0.00s)
    --- PASS: TestSerialisedSizes/fields,15,protobuf,size,2601 (0.00s)
    --- PASS: TestSerialisedSizes/fields,15,json,size,3389 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,0,protobuf,size,81 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,0,json,size,248 (0.00s)
    --- PASS: TestSerialisedSizes/2kb_msg,protobuf,size,2132 (0.00s)
    --- PASS: TestSerialisedSizes/2kb_msg,json,size,2684 (0.00s)
    --- PASS: TestSerialisedSizes/fields,1,protobuf,size,249 (0.00s)
    --- PASS: TestSerialisedSizes/fields,1,json,size,467 (0.00s)
    --- PASS: TestSerialisedSizes/fields,3,protobuf,size,585 (0.00s)
    --- PASS: TestSerialisedSizes/fields,3,json,size,912 (0.00s)
    --- PASS: TestSerialisedSizes/fields,127,protobuf,size,21417 (0.00s)
    --- PASS: TestSerialisedSizes/fields,127,json,size,26916 (0.00s)
    --- PASS: TestSerialisedSizes/fields,1023,protobuf,size,171945 (0.00s)
    --- PASS: TestSerialisedSizes/fields,1023,json,size,216650 (0.00s)
    --- PASS: TestSerialisedSizes/small_msg,protobuf,size,106 (0.00s)
    --- PASS: TestSerialisedSizes/small_msg,json,size,271 (0.00s)
    --- PASS: TestSerialisedSizes/fields,63,protobuf,size,10665 (0.00s)
    --- PASS: TestSerialisedSizes/fields,63,json,size,13704 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,127,protobuf,size,6050 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,127,json,size,11231 (0.00s)
    --- PASS: TestSerialisedSizes/16kb_msg,protobuf,size,16469 (0.00s)
    --- PASS: TestSerialisedSizes/16kb_msg,json,size,19729 (0.00s)
    --- PASS: TestSerialisedSizes/fields,31,protobuf,size,5289 (0.00s)
    --- PASS: TestSerialisedSizes/fields,31,json,size,6730 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,7,protobuf,size,410 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,7,json,size,852 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,63,protobuf,size,3042 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,63,json,size,5784 (0.00s)
    --- PASS: TestSerialisedSizes/no_msg,protobuf,size,81 (0.00s)
    --- PASS: TestSerialisedSizes/no_msg,json,size,248 (0.00s)
    --- PASS: TestSerialisedSizes/fields,7,protobuf,size,1257 (0.00s)
    --- PASS: TestSerialisedSizes/fields,7,json,size,1724 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,3,protobuf,size,222 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,3,json,size,513 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,31,protobuf,size,1538 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,31,json,size,2898 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,255,protobuf,size,12066 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,255,json,size,22317 (0.00s)
    --- PASS: TestSerialisedSizes/8kb_msg,protobuf,size,8276 (0.00s)
    --- PASS: TestSerialisedSizes/8kb_msg,json,size,9921 (0.00s)
    --- PASS: TestSerialisedSizes/fields,255,protobuf,size,42921 (0.00s)
    --- PASS: TestSerialisedSizes/fields,255,json,size,54159 (0.00s)
    --- PASS: TestSerialisedSizes/fields,511,protobuf,size,85929 (0.00s)
    --- PASS: TestSerialisedSizes/fields,511,json,size,108538 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,1,protobuf,size,128 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,1,json,size,344 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,511,protobuf,size,24098 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,511,json,size,44621 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,1023,protobuf,size,48162 (0.00s)
    --- PASS: TestSerialisedSizes/metrics,1023,json,size,89023 (0.00s)
    --- PASS: TestSerialisedSizes/4kb_msg,protobuf,size,4180 (0.00s)
    --- PASS: TestSerialisedSizes/4kb_msg,json,size,5154 (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/Scout24/protobuf-vs-json
BenchmarkIt/serialise,protobuf,fields,1023-8         	   10000	    207737 ns/op	8276216.81 MB/s	  172015 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,1023-8             	    2000	    682502 ns/op	635231.60 MB/s	  221690 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,1023-8             	    5000	    296530 ns/op	 579.86 MB/s	  245814 B/op	    3085 allocs/op
BenchmarkIt/parse,json,fields,1023-8                 	     500	   2927776 ns/op	  74.08 MB/s	  494688 B/op	    3994 allocs/op
BenchmarkIt/parse,json-validate,fields,1023-8        	     200	   8577722 ns/op	  25.28 MB/s	 1765098 B/op	   39524 allocs/op
BenchmarkIt/serialise,protobuf,metrics,1-8           	 3000000	       454 ns/op	844965085.50 MB/s	     127 B/op	       1 allocs/op
BenchmarkIt/serialise,json,metrics,1-8               	 1000000	      2074 ns/op	163866272.80 MB/s	     496 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,1-8               	 2000000	       602 ns/op	 212.31 MB/s	     439 B/op	       8 allocs/op
BenchmarkIt/parse,json,metrics,1-8                   	  200000	      7131 ns/op	  47.68 MB/s	     903 B/op	      18 allocs/op
BenchmarkIt/parse,json-validate,metrics,1-8          	   50000	     34055 ns/op	   9.98 MB/s	    6744 B/op	     207 allocs/op
BenchmarkIt/serialise,protobuf,metrics,7-8           	 1000000	      1095 ns/op	374170472.13 MB/s	     415 B/op	       1 allocs/op
BenchmarkIt/serialise,json,metrics,7-8               	  300000	      4417 ns/op	58273427.87 MB/s	    1040 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,7-8               	 1000000	      1745 ns/op	 234.91 MB/s	    1127 B/op	      23 allocs/op
BenchmarkIt/parse,json,metrics,7-8                   	  100000	     15853 ns/op	  54.12 MB/s	    2167 B/op	      45 allocs/op
BenchmarkIt/parse,json-validate,metrics,7-8          	   20000	     83282 ns/op	  10.30 MB/s	   18001 B/op	     584 allocs/op
BenchmarkIt/serialise,protobuf,metrics,15-8          	 1000000	      1945 ns/op	404061967.88 MB/s	     895 B/op	       1 allocs/op
BenchmarkIt/serialise,json,metrics,15-8              	  200000	      7440 ns/op	41744836.44 MB/s	    1936 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,15-8              	  500000	      3066 ns/op	 256.28 MB/s	    2023 B/op	      40 allocs/op
BenchmarkIt/parse,json,metrics,15-8                  	   50000	     28541 ns/op	  54.41 MB/s	    4415 B/op	      81 allocs/op
BenchmarkIt/parse,json-validate,metrics,15-8         	   10000	    149407 ns/op	  10.39 MB/s	   33695 B/op	    1085 allocs/op
BenchmarkIt/serialise,protobuf,bytes,3-8             	  500000	      2715 ns/op	581132294.54 MB/s	    3199 B/op	       1 allocs/op
BenchmarkIt/serialise,json,bytes,3-8                 	  200000	      9230 ns/op	84396061.61 MB/s	    4241 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,bytes,3-8                 	  500000	      2737 ns/op	1152.71 MB/s	    3407 B/op	       6 allocs/op
BenchmarkIt/parse,json,bytes,3-8                     	   30000	     40759 ns/op	  95.56 MB/s	    7759 B/op	      14 allocs/op
BenchmarkIt/parse,json-validate,bytes,3-8            	   20000	    100458 ns/op	  38.77 MB/s	   19085 B/op	     146 allocs/op
BenchmarkIt/serialise,protobuf,fields,255-8          	   30000	     51995 ns/op	24763195.46 MB/s	   49150 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,255-8              	   10000	    171074 ns/op	3144457.70 MB/s	   57530 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,255-8              	   20000	     72298 ns/op	 593.66 MB/s	   61540 B/op	     779 allocs/op
BenchmarkIt/parse,json,fields,255-8                  	    2000	    725248 ns/op	  74.18 MB/s	  126448 B/op	    1019 allocs/op
BenchmarkIt/parse,json-validate,fields,255-8         	    1000	   2162408 ns/op	  24.88 MB/s	  444464 B/op	    9937 allocs/op
BenchmarkIt/serialise,protobuf,metrics,511-8         	   30000	     53131 ns/op	13606285.71 MB/s	   24575 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,511-8             	   10000	    209104 ns/op	2140629.93 MB/s	   49321 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,511-8             	   20000	     83754 ns/op	 287.72 MB/s	   57573 B/op	    1037 allocs/op
BenchmarkIt/parse,json,metrics,511-8                 	    2000	    826765 ns/op	  54.15 MB/s	  140770 B/op	    2018 allocs/op
BenchmarkIt/parse,json-validate,metrics,511-8        	     300	   4261539 ns/op	  10.50 MB/s	 1023162 B/op	   32038 allocs/op
BenchmarkIt/serialise,protobuf,bytes,0-8             	 5000000	       331 ns/op	1221834838.42 MB/s	      96 B/op	       1 allocs/op
BenchmarkIt/serialise,json,bytes,0-8                 	 1000000	      1745 ns/op	141472311.40 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,bytes,0-8                 	 3000000	       405 ns/op	 199.66 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/parse,json,bytes,0-8                     	  300000	      4875 ns/op	  50.66 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/parse,json-validate,bytes,0-8            	   50000	     23965 ns/op	  10.31 MB/s	    4723 B/op	     141 allocs/op
BenchmarkIt/serialise,protobuf,metrics,0-8           	 5000000	       326 ns/op	1241967854.19 MB/s	      96 B/op	       1 allocs/op
BenchmarkIt/serialise,json,metrics,0-8               	 1000000	      1601 ns/op	154884572.64 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,0-8               	 3000000	       405 ns/op	 199.60 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/parse,json,metrics,0-8                   	  300000	      4857 ns/op	  51.06 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/parse,json-validate,metrics,0-8          	   50000	     24152 ns/op	  10.27 MB/s	    4722 B/op	     141 allocs/op
BenchmarkIt/serialise,protobuf,fields,0-8            	 5000000	       323 ns/op	1250735278.08 MB/s	      95 B/op	       0 allocs/op
BenchmarkIt/serialise,json,fields,0-8                	 1000000	      1598 ns/op	155191473.78 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,0-8                	 5000000	       397 ns/op	 203.69 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/parse,json,fields,0-8                    	  300000	      4849 ns/op	  51.14 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/parse,json-validate,fields,0-8           	   50000	     24005 ns/op	  10.33 MB/s	    4724 B/op	     141 allocs/op
BenchmarkIt/serialise,protobuf,fields,7-8            	 1000000	      1834 ns/op	685055862.66 MB/s	    1279 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,7-8                	  300000	      5055 ns/op	103134745.95 MB/s	    1936 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,7-8                	  500000	      2635 ns/op	 476.92 MB/s	    2023 B/op	      30 allocs/op
BenchmarkIt/parse,json,fields,7-8                    	   50000	     24352 ns/op	  71.37 MB/s	    3919 B/op	      45 allocs/op
BenchmarkIt/parse,json-validate,fields,7-8           	   20000	     81512 ns/op	  21.32 MB/s	   16510 B/op	     416 allocs/op
BenchmarkIt/serialise,protobuf,fields,15-8           	  500000	      3518 ns/op	369651422.92 MB/s	    2687 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,15-8               	  200000	      8810 ns/op	76978799.49 MB/s	    3601 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,15-8               	  300000	      4967 ns/op	 523.59 MB/s	    3943 B/op	      55 allocs/op
BenchmarkIt/parse,json,fields,15-8                   	   30000	     45520 ns/op	  74.49 MB/s	    7967 B/op	      80 allocs/op
BenchmarkIt/parse,json-validate,fields,15-8          	   10000	    150443 ns/op	  22.54 MB/s	   30168 B/op	     723 allocs/op
BenchmarkIt/serialise,protobuf,metrics,3-8           	 2000000	       659 ns/op	673593443.91 MB/s	     223 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,3-8               	  500000	      2806 ns/op	90321038.10 MB/s	     656 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,3-8               	 1000000	      1021 ns/op	 217.42 MB/s	     679 B/op	      14 allocs/op
BenchmarkIt/parse,json,metrics,3-8                   	  200000	      9445 ns/op	  53.68 MB/s	    1087 B/op	      26 allocs/op
BenchmarkIt/parse,json-validate,metrics,3-8          	   30000	     49787 ns/op	  10.18 MB/s	   10277 B/op	     333 allocs/op
BenchmarkIt/serialise,protobuf,metrics,31-8          	  500000	      3608 ns/op	213121649.15 MB/s	    1791 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,31-8              	  100000	     13372 ns/op	21977169.63 MB/s	    3217 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,31-8              	  300000	      5659 ns/op	 271.77 MB/s	    3815 B/op	      73 allocs/op
BenchmarkIt/parse,json,metrics,31-8                  	   30000	     53251 ns/op	  55.19 MB/s	    8967 B/op	     146 allocs/op
BenchmarkIt/parse,json-validate,metrics,31-8         	    5000	    278645 ns/op	  10.55 MB/s	   65089 B/op	    2075 allocs/op
BenchmarkIt/serialise,protobuf,metrics,127-8         	  100000	     13543 ns/op	44670221.19 MB/s	    6143 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,127-8             	   30000	     51146 ns/op	6636551.87 MB/s	   12438 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,127-8             	  100000	     20543 ns/op	 294.50 MB/s	   14567 B/op	     267 allocs/op
BenchmarkIt/parse,json,metrics,127-8                 	   10000	    200366 ns/op	  56.47 MB/s	   31476 B/op	     521 allocs/op
BenchmarkIt/parse,json-validate,metrics,127-8        	    2000	   1034874 ns/op	  10.93 MB/s	  250334 B/op	    8030 allocs/op
BenchmarkIt/serialise,protobuf,bytes,15-8            	  100000	     11948 ns/op	129248222.16 MB/s	   16383 B/op	       0 allocs/op
BenchmarkIt/serialise,json,bytes,15-8                	   30000	     48128 ns/op	11443438.31 MB/s	   18584 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,bytes,15-8                	  100000	     12147 ns/op	1271.38 MB/s	   16719 B/op	       6 allocs/op
BenchmarkIt/parse,json,bytes,15-8                    	   10000	    189674 ns/op	  96.79 MB/s	   35404 B/op	      14 allocs/op
BenchmarkIt/parse,json-validate,bytes,15-8           	    3000	    399336 ns/op	  45.97 MB/s	   74406 B/op	     146 allocs/op
BenchmarkIt/serialise,protobuf,fields,1-8            	 3000000	       538 ns/op	1386267957.47 MB/s	     255 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,1-8                	 1000000	      2137 ns/op	217030717.03 MB/s	     624 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,1-8                	 2000000	       706 ns/op	 352.58 MB/s	     567 B/op	       9 allocs/op
BenchmarkIt/parse,json,fields,1-8                    	  200000	      7450 ns/op	  62.28 MB/s	    1135 B/op	      18 allocs/op
BenchmarkIt/parse,json-validate,fields,1-8           	   50000	     32326 ns/op	  14.35 MB/s	    6503 B/op	     183 allocs/op
BenchmarkIt/serialise,protobuf,fields,31-8           	  200000	      6714 ns/op	157540766.60 MB/s	    5375 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,31-8               	  100000	     17314 ns/op	38971729.78 MB/s	    6930 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,31-8               	  200000	      9374 ns/op	 564.19 MB/s	    7783 B/op	     104 allocs/op
BenchmarkIt/parse,json,fields,31-8                   	   20000	     89444 ns/op	  75.44 MB/s	   16287 B/op	     147 allocs/op
BenchmarkIt/parse,json-validate,fields,31-8          	    5000	    277949 ns/op	  24.28 MB/s	   57892 B/op	    1333 allocs/op
BenchmarkIt/serialise,protobuf,fields,63-8           	  100000	     12860 ns/op	82927475.61 MB/s	   10879 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,63-8               	   50000	     37730 ns/op	18010114.55 MB/s	   14485 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,63-8               	  100000	     18450 ns/op	 578.04 MB/s	   15463 B/op	     201 allocs/op
BenchmarkIt/parse,json,fields,63-8                   	   10000	    180163 ns/op	  75.44 MB/s	   30316 B/op	     270 allocs/op
BenchmarkIt/parse,json-validate,fields,63-8          	    3000	    531291 ns/op	  25.58 MB/s	  110774 B/op	    2538 allocs/op
BenchmarkIt/serialise,protobuf,metrics,63-8          	  200000	      6855 ns/op	88743093.06 MB/s	    3071 B/op	       1 allocs/op
BenchmarkIt/serialise,json,metrics,63-8              	   50000	     25244 ns/op	11328998.40 MB/s	    6290 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,63-8              	  200000	     10365 ns/op	 293.48 MB/s	    7399 B/op	     138 allocs/op
BenchmarkIt/parse,json,metrics,63-8                  	   20000	     98383 ns/op	  58.14 MB/s	   14767 B/op	     270 allocs/op
BenchmarkIt/parse,json-validate,metrics,63-8         	    3000	    517868 ns/op	  11.05 MB/s	  124569 B/op	    4050 allocs/op
BenchmarkIt/serialise,protobuf,fields,511-8          	   10000	    100446 ns/op	8553871.38 MB/s	   90102 B/op	       0 allocs/op
BenchmarkIt/serialise,json,fields,511-8              	    5000	    329254 ns/op	1643947.72 MB/s	  114909 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,511-8              	   10000	    142537 ns/op	 602.85 MB/s	  122971 B/op	    1548 allocs/op
BenchmarkIt/parse,json,fields,511-8                  	    1000	   1444518 ns/op	  74.96 MB/s	  263704 B/op	    2025 allocs/op
BenchmarkIt/parse,json-validate,fields,511-8         	     300	   4065238 ns/op	  26.63 MB/s	  899939 B/op	   19831 allocs/op
BenchmarkIt/serialise,protobuf,bytes,7-8             	  200000	      5955 ns/op	243546081.25 MB/s	    8191 B/op	       1 allocs/op
BenchmarkIt/serialise,json,bytes,7-8                 	  100000	     20106 ns/op	43244495.56 MB/s	    9619 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,bytes,7-8                 	  200000	      6079 ns/op	1192.78 MB/s	    8527 B/op	       6 allocs/op
BenchmarkIt/parse,json,bytes,7-8                     	   20000	     87894 ns/op	  98.93 MB/s	   18255 B/op	      14 allocs/op
BenchmarkIt/parse,json-validate,bytes,7-8            	   10000	    195532 ns/op	  44.47 MB/s	   40092 B/op	     146 allocs/op
BenchmarkIt/serialise,protobuf,fields,3-8            	 2000000	       968 ns/op	1208589409.56 MB/s	     639 B/op	       0 allocs/op
BenchmarkIt/serialise,json,fields,3-8                	  500000	      2991 ns/op	146566178.24 MB/s	    1040 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,3-8                	 1000000	      1370 ns/op	 426.83 MB/s	    1063 B/op	      17 allocs/op
BenchmarkIt/parse,json,fields,3-8                    	  100000	     12242 ns/op	  71.64 MB/s	    1887 B/op	      26 allocs/op
BenchmarkIt/parse,json-validate,fields,3-8           	   30000	     47263 ns/op	  18.56 MB/s	    9689 B/op	     261 allocs/op
BenchmarkIt/serialise,protobuf,fields,127-8          	   50000	     25264 ns/op	42384340.32 MB/s	   21759 B/op	       1 allocs/op
BenchmarkIt/serialise,json,fields,127-8              	   20000	     78722 ns/op	6837824.74 MB/s	   27424 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,fields,127-8              	   50000	     35758 ns/op	 598.93 MB/s	   30823 B/op	     394 allocs/op
BenchmarkIt/parse,json,fields,127-8                  	    5000	    353503 ns/op	  76.14 MB/s	   62035 B/op	     523 allocs/op
BenchmarkIt/parse,json-validate,fields,127-8         	    2000	   1029169 ns/op	  26.15 MB/s	  220934 B/op	    4987 allocs/op
BenchmarkIt/serialise,protobuf,metrics,255-8         	   50000	     26706 ns/op	22589248.74 MB/s	   12287 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,255-8             	   20000	     98569 ns/op	4535861.61 MB/s	   24734 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,255-8             	   50000	     39818 ns/op	 303.03 MB/s	   28903 B/op	     524 allocs/op
BenchmarkIt/parse,json,metrics,255-8                 	    3000	    387255 ns/op	  57.73 MB/s	   66953 B/op	    1017 allocs/op
BenchmarkIt/parse,json-validate,metrics,255-8        	    1000	   1996467 ns/op	  11.20 MB/s	  507325 B/op	   16048 allocs/op
BenchmarkIt/serialise,protobuf,metrics,1023-8        	   10000	    106817 ns/op	4508345.52 MB/s	   49147 B/op	       0 allocs/op
BenchmarkIt/serialise,json,metrics,1023-8            	    3000	    405982 ns/op	659453.59 MB/s	   90353 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,metrics,1023-8            	   10000	    156853 ns/op	 307.05 MB/s	  114908 B/op	    2062 allocs/op
BenchmarkIt/parse,json,metrics,1023-8                	    1000	   1529810 ns/op	  58.35 MB/s	  241311 B/op	    4016 allocs/op
BenchmarkIt/parse,json-validate,metrics,1023-8       	     200	   8083167 ns/op	  11.04 MB/s	 2007211 B/op	   63992 allocs/op
BenchmarkIt/serialise,protobuf,bytes,1-8             	 1000000	      1118 ns/op	990981720.14 MB/s	    1151 B/op	       0 allocs/op
BenchmarkIt/serialise,json,bytes,1-8                 	  300000	      4082 ns/op	106626828.01 MB/s	    1680 B/op	       4 allocs/op
BenchmarkIt/parse,protobuf,bytes,1-8                 	 1000000	      1213 ns/op	 912.71 MB/s	    1359 B/op	       6 allocs/op
BenchmarkIt/parse,json,bytes,1-8                     	  100000	     16079 ns/op	  90.24 MB/s	    2895 B/op	      14 allocs/op
BenchmarkIt/parse,json-validate,bytes,1-8            	   30000	     46773 ns/op	  31.02 MB/s	    9350 B/op	     146 allocs/op
PASS
ok  	github.com/Scout24/protobuf-vs-json	242.882s
```
