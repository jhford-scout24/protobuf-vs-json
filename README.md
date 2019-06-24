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
=== RUN   TestSerialisedSizes/7_metrics_PROTOBUF:_410_bytes
=== RUN   TestSerialisedSizes/7_metrics_JSON:_____865_bytes
=== RUN   TestSerialisedSizes/31_metrics_PROTOBUF:_1538_bytes
=== RUN   TestSerialisedSizes/31_metrics_JSON:_____2946_bytes
=== RUN   TestSerialisedSizes/255_metrics_PROTOBUF:_12066_bytes
=== RUN   TestSerialisedSizes/255_metrics_JSON:_____22408_bytes
=== RUN   TestSerialisedSizes/511_metrics_PROTOBUF:_24098_bytes
=== RUN   TestSerialisedSizes/511_metrics_JSON:_____44786_bytes
=== RUN   TestSerialisedSizes/0_fields_PROTOBUF:_81_bytes
=== RUN   TestSerialisedSizes/0_fields_JSON:_____247_bytes
=== RUN   TestSerialisedSizes/7_fields_PROTOBUF:_1257_bytes
=== RUN   TestSerialisedSizes/7_fields_JSON:_____1678_bytes
=== RUN   TestSerialisedSizes/127_metrics_PROTOBUF:_6050_bytes
=== RUN   TestSerialisedSizes/127_metrics_JSON:_____11222_bytes
=== RUN   TestSerialisedSizes/0_metrics_PROTOBUF:_81_bytes
=== RUN   TestSerialisedSizes/0_metrics_JSON:_____247_bytes
=== RUN   TestSerialisedSizes/63_metrics_PROTOBUF:_3042_bytes
=== RUN   TestSerialisedSizes/63_metrics_JSON:_____5754_bytes
=== RUN   TestSerialisedSizes/small_msg_PROTOBUF:_106_bytes
=== RUN   TestSerialisedSizes/small_msg_JSON:_____270_bytes
=== RUN   TestSerialisedSizes/15_fields_PROTOBUF:_2601_bytes
=== RUN   TestSerialisedSizes/15_fields_JSON:_____3429_bytes
=== RUN   TestSerialisedSizes/63_fields_PROTOBUF:_10665_bytes
=== RUN   TestSerialisedSizes/63_fields_JSON:_____13487_bytes
=== RUN   TestSerialisedSizes/1023_metrics_PROTOBUF:_48162_bytes
=== RUN   TestSerialisedSizes/1023_metrics_JSON:_____88972_bytes
=== RUN   TestSerialisedSizes/3_fields_PROTOBUF:_585_bytes
=== RUN   TestSerialisedSizes/3_fields_JSON:_____867_bytes
=== RUN   TestSerialisedSizes/31_fields_PROTOBUF:_5289_bytes
=== RUN   TestSerialisedSizes/31_fields_JSON:_____6871_bytes
=== RUN   TestSerialisedSizes/127_fields_PROTOBUF:_21417_bytes
=== RUN   TestSerialisedSizes/127_fields_JSON:_____27079_bytes
=== RUN   TestSerialisedSizes/511_fields_PROTOBUF:_85929_bytes
=== RUN   TestSerialisedSizes/511_fields_JSON:_____108505_bytes
=== RUN   TestSerialisedSizes/no_msg_PROTOBUF:_81_bytes
=== RUN   TestSerialisedSizes/no_msg_JSON:_____248_bytes
=== RUN   TestSerialisedSizes/2kb_msg_PROTOBUF:_2132_bytes
=== RUN   TestSerialisedSizes/2kb_msg_JSON:_____2706_bytes
=== RUN   TestSerialisedSizes/8kb_msg_PROTOBUF:_8276_bytes
=== RUN   TestSerialisedSizes/8kb_msg_JSON:_____10078_bytes
=== RUN   TestSerialisedSizes/4kb_msg_PROTOBUF:_4180_bytes
=== RUN   TestSerialisedSizes/4kb_msg_JSON:_____5082_bytes
=== RUN   TestSerialisedSizes/1_fields_PROTOBUF:_249_bytes
=== RUN   TestSerialisedSizes/1_fields_JSON:_____478_bytes
=== RUN   TestSerialisedSizes/1023_fields_PROTOBUF:_171945_bytes
=== RUN   TestSerialisedSizes/1023_fields_JSON:_____216064_bytes
=== RUN   TestSerialisedSizes/1_metrics_PROTOBUF:_128_bytes
=== RUN   TestSerialisedSizes/1_metrics_JSON:_____330_bytes
=== RUN   TestSerialisedSizes/15_metrics_PROTOBUF:_786_bytes
=== RUN   TestSerialisedSizes/15_metrics_JSON:_____1551_bytes
=== RUN   TestSerialisedSizes/16kb_msg_PROTOBUF:_16469_bytes
=== RUN   TestSerialisedSizes/16kb_msg_JSON:_____19773_bytes
=== RUN   TestSerialisedSizes/255_fields_PROTOBUF:_42921_bytes
=== RUN   TestSerialisedSizes/255_fields_JSON:_____54271_bytes
=== RUN   TestSerialisedSizes/3_metrics_PROTOBUF:_222_bytes
=== RUN   TestSerialisedSizes/3_metrics_JSON:_____502_bytes
--- PASS: TestSerialisedSizes (0.01s)
    --- PASS: TestSerialisedSizes/7_metrics_PROTOBUF:_410_bytes (0.00s)
    --- PASS: TestSerialisedSizes/7_metrics_JSON:_____865_bytes (0.00s)
    --- PASS: TestSerialisedSizes/31_metrics_PROTOBUF:_1538_bytes (0.00s)
    --- PASS: TestSerialisedSizes/31_metrics_JSON:_____2946_bytes (0.00s)
    --- PASS: TestSerialisedSizes/255_metrics_PROTOBUF:_12066_bytes (0.00s)
    --- PASS: TestSerialisedSizes/255_metrics_JSON:_____22408_bytes (0.00s)
    --- PASS: TestSerialisedSizes/511_metrics_PROTOBUF:_24098_bytes (0.00s)
    --- PASS: TestSerialisedSizes/511_metrics_JSON:_____44786_bytes (0.00s)
    --- PASS: TestSerialisedSizes/0_fields_PROTOBUF:_81_bytes (0.00s)
    --- PASS: TestSerialisedSizes/0_fields_JSON:_____247_bytes (0.00s)
    --- PASS: TestSerialisedSizes/7_fields_PROTOBUF:_1257_bytes (0.00s)
    --- PASS: TestSerialisedSizes/7_fields_JSON:_____1678_bytes (0.00s)
    --- PASS: TestSerialisedSizes/127_metrics_PROTOBUF:_6050_bytes (0.00s)
    --- PASS: TestSerialisedSizes/127_metrics_JSON:_____11222_bytes (0.00s)
    --- PASS: TestSerialisedSizes/0_metrics_PROTOBUF:_81_bytes (0.00s)
    --- PASS: TestSerialisedSizes/0_metrics_JSON:_____247_bytes (0.00s)
    --- PASS: TestSerialisedSizes/63_metrics_PROTOBUF:_3042_bytes (0.00s)
    --- PASS: TestSerialisedSizes/63_metrics_JSON:_____5754_bytes (0.00s)
    --- PASS: TestSerialisedSizes/small_msg_PROTOBUF:_106_bytes (0.00s)
    --- PASS: TestSerialisedSizes/small_msg_JSON:_____270_bytes (0.00s)
    --- PASS: TestSerialisedSizes/15_fields_PROTOBUF:_2601_bytes (0.00s)
    --- PASS: TestSerialisedSizes/15_fields_JSON:_____3429_bytes (0.00s)
    --- PASS: TestSerialisedSizes/63_fields_PROTOBUF:_10665_bytes (0.00s)
    --- PASS: TestSerialisedSizes/63_fields_JSON:_____13487_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1023_metrics_PROTOBUF:_48162_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1023_metrics_JSON:_____88972_bytes (0.00s)
    --- PASS: TestSerialisedSizes/3_fields_PROTOBUF:_585_bytes (0.00s)
    --- PASS: TestSerialisedSizes/3_fields_JSON:_____867_bytes (0.00s)
    --- PASS: TestSerialisedSizes/31_fields_PROTOBUF:_5289_bytes (0.00s)
    --- PASS: TestSerialisedSizes/31_fields_JSON:_____6871_bytes (0.00s)
    --- PASS: TestSerialisedSizes/127_fields_PROTOBUF:_21417_bytes (0.00s)
    --- PASS: TestSerialisedSizes/127_fields_JSON:_____27079_bytes (0.00s)
    --- PASS: TestSerialisedSizes/511_fields_PROTOBUF:_85929_bytes (0.00s)
    --- PASS: TestSerialisedSizes/511_fields_JSON:_____108505_bytes (0.00s)
    --- PASS: TestSerialisedSizes/no_msg_PROTOBUF:_81_bytes (0.00s)
    --- PASS: TestSerialisedSizes/no_msg_JSON:_____248_bytes (0.00s)
    --- PASS: TestSerialisedSizes/2kb_msg_PROTOBUF:_2132_bytes (0.00s)
    --- PASS: TestSerialisedSizes/2kb_msg_JSON:_____2706_bytes (0.00s)
    --- PASS: TestSerialisedSizes/8kb_msg_PROTOBUF:_8276_bytes (0.00s)
    --- PASS: TestSerialisedSizes/8kb_msg_JSON:_____10078_bytes (0.00s)
    --- PASS: TestSerialisedSizes/4kb_msg_PROTOBUF:_4180_bytes (0.00s)
    --- PASS: TestSerialisedSizes/4kb_msg_JSON:_____5082_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1_fields_PROTOBUF:_249_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1_fields_JSON:_____478_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1023_fields_PROTOBUF:_171945_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1023_fields_JSON:_____216064_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1_metrics_PROTOBUF:_128_bytes (0.00s)
    --- PASS: TestSerialisedSizes/1_metrics_JSON:_____330_bytes (0.00s)
    --- PASS: TestSerialisedSizes/15_metrics_PROTOBUF:_786_bytes (0.00s)
    --- PASS: TestSerialisedSizes/15_metrics_JSON:_____1551_bytes (0.00s)
    --- PASS: TestSerialisedSizes/16kb_msg_PROTOBUF:_16469_bytes (0.00s)
    --- PASS: TestSerialisedSizes/16kb_msg_JSON:_____19773_bytes (0.00s)
    --- PASS: TestSerialisedSizes/255_fields_PROTOBUF:_42921_bytes (0.00s)
    --- PASS: TestSerialisedSizes/255_fields_JSON:_____54271_bytes (0.00s)
    --- PASS: TestSerialisedSizes/3_metrics_PROTOBUF:_222_bytes (0.00s)
    --- PASS: TestSerialisedSizes/3_metrics_JSON:_____502_bytes (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/Scout24/protobuf-vs-json
BenchmarkIt/4kb_msg_PROTOBUF_SERIALISE-8         	  500000	      3488 ns/op	599048786.36 MB/s	    4863 B/op	       1 allocs/op
BenchmarkIt/4kb_msg_JSON_SERIALISE-8             	  100000	     11997 ns/op	42934855.01 MB/s	    5521 B/op	       4 allocs/op
BenchmarkIt/4kb_msg_PROTOBUF_PARSE-8             	  500000	      3576 ns/op	1168.70 MB/s	    4431 B/op	       6 allocs/op
BenchmarkIt/4kb_msg_JSON_PARSE-8                 	   30000	     54119 ns/op	  95.18 MB/s	   10063 B/op	      14 allocs/op
BenchmarkIt/4kb_msg_JSON_VALIDATE-PARSE-8        	   10000	    127927 ns/op	  40.27 MB/s	   23697 B/op	     146 allocs/op
BenchmarkIt/1_fields_PROTOBUF_SERIALISE-8        	 3000000	       538 ns/op	1386884524.25 MB/s	     255 B/op	       0 allocs/op
BenchmarkIt/1_fields_JSON_SERIALISE-8            	  500000	      2035 ns/op	109295610.13 MB/s	     592 B/op	       4 allocs/op
BenchmarkIt/1_fields_PROTOBUF_PARSE-8            	 2000000	       710 ns/op	 350.32 MB/s	     567 B/op	       9 allocs/op
BenchmarkIt/1_fields_JSON_PARSE-8                	  200000	      7202 ns/op	  61.78 MB/s	    1119 B/op	      18 allocs/op
BenchmarkIt/1_fields_JSON_VALIDATE-PARSE-8       	   50000	     32775 ns/op	  13.58 MB/s	    6470 B/op	     183 allocs/op
BenchmarkIt/7_fields_PROTOBUF_SERIALISE-8        	 1000000	      1813 ns/op	693075638.19 MB/s	    1279 B/op	       1 allocs/op
BenchmarkIt/7_fields_JSON_SERIALISE-8            	  300000	      5048 ns/op	105703759.06 MB/s	    1936 B/op	       4 allocs/op
BenchmarkIt/7_fields_PROTOBUF_PARSE-8            	  500000	      2574 ns/op	 488.22 MB/s	    2023 B/op	      30 allocs/op
BenchmarkIt/7_fields_JSON_PARSE-8                	  100000	     23437 ns/op	  75.90 MB/s	    3903 B/op	      44 allocs/op
BenchmarkIt/7_fields_JSON_VALIDATE-PARSE-8       	   20000	     80021 ns/op	  22.23 MB/s	   16478 B/op	     414 allocs/op
BenchmarkIt/1023_metrics_PROTOBUF_SERIALISE-8    	   10000	    103980 ns/op	4631348.17 MB/s	   49147 B/op	       1 allocs/op
BenchmarkIt/1023_metrics_JSON_SERIALISE-8        	    3000	    399114 ns/op	669389.25 MB/s	   90362 B/op	       4 allocs/op
BenchmarkIt/1023_metrics_PROTOBUF_PARSE-8        	   10000	    155728 ns/op	 309.27 MB/s	  114908 B/op	    2062 allocs/op
BenchmarkIt/1023_metrics_JSON_PARSE-8            	    1000	   1540636 ns/op	  57.82 MB/s	  240000 B/op	    3991 allocs/op
BenchmarkIt/1023_metrics_JSON_VALIDATE-PARSE-8   	     200	   8008926 ns/op	  11.12 MB/s	 2004584 B/op	   63940 allocs/op
BenchmarkIt/16kb_msg_PROTOBUF_SERIALISE-8        	  100000	     12117 ns/op	135914277.19 MB/s	   18431 B/op	       0 allocs/op
BenchmarkIt/16kb_msg_JSON_SERIALISE-8            	   30000	     47118 ns/op	12437984.70 MB/s	   20631 B/op	       4 allocs/op
BenchmarkIt/16kb_msg_PROTOBUF_PARSE-8            	  100000	     12346 ns/op	1333.94 MB/s	   16719 B/op	       7 allocs/op
BenchmarkIt/16kb_msg_JSON_PARSE-8                	   10000	    197123 ns/op	  99.11 MB/s	   37452 B/op	      14 allocs/op
BenchmarkIt/16kb_msg_JSON_VALIDATE-PARSE-8       	    3000	    411678 ns/op	  47.45 MB/s	   78499 B/op	     146 allocs/op
BenchmarkIt/31_fields_PROTOBUF_SERIALISE-8       	  200000	      6600 ns/op	160259482.00 MB/s	    5375 B/op	       1 allocs/op
BenchmarkIt/31_fields_JSON_SERIALISE-8           	  100000	     17904 ns/op	38039575.74 MB/s	    7058 B/op	       4 allocs/op
BenchmarkIt/31_fields_PROTOBUF_PARSE-8           	  200000	      9393 ns/op	 563.04 MB/s	    7783 B/op	     104 allocs/op
BenchmarkIt/31_fields_JSON_PARSE-8               	   20000	     90769 ns/op	  75.04 MB/s	   16287 B/op	     144 allocs/op
BenchmarkIt/31_fields_JSON_VALIDATE-PARSE-8      	    5000	    277280 ns/op	  24.56 MB/s	   57894 B/op	    1327 allocs/op
BenchmarkIt/1023_fields_PROTOBUF_SERIALISE-8     	   10000	    198534 ns/op	8659827.69 MB/s	  172014 B/op	       1 allocs/op
BenchmarkIt/1023_fields_JSON_SERIALISE-8         	    2000	    649742 ns/op	667183.25 MB/s	  221618 B/op	       4 allocs/op
BenchmarkIt/1023_fields_PROTOBUF_PARSE-8         	    5000	    284744 ns/op	 603.86 MB/s	  245814 B/op	    3085 allocs/op
BenchmarkIt/1023_fields_JSON_PARSE-8             	     500	   2783351 ns/op	  77.91 MB/s	  494225 B/op	    3980 allocs/op
BenchmarkIt/1023_fields_JSON_VALIDATE-PARSE-8    	     200	   8084212 ns/op	  26.82 MB/s	 1764168 B/op	   39494 allocs/op
BenchmarkIt/3_metrics_PROTOBUF_SERIALISE-8       	 2000000	       641 ns/op	692542000.55 MB/s	     223 B/op	       1 allocs/op
BenchmarkIt/3_metrics_JSON_SERIALISE-8           	  500000	      2754 ns/op	93489650.68 MB/s	     720 B/op	       4 allocs/op
BenchmarkIt/3_metrics_PROTOBUF_PARSE-8           	 1000000	      1009 ns/op	 219.93 MB/s	     679 B/op	      14 allocs/op
BenchmarkIt/3_metrics_JSON_PARSE-8               	  200000	      9282 ns/op	  55.48 MB/s	    1119 B/op	      26 allocs/op
BenchmarkIt/3_metrics_JSON_VALIDATE-PARSE-8      	   30000	     48536 ns/op	  10.61 MB/s	   10341 B/op	     333 allocs/op
BenchmarkIt/127_metrics_PROTOBUF_SERIALISE-8     	  100000	     13406 ns/op	45125825.45 MB/s	    6143 B/op	       1 allocs/op
BenchmarkIt/127_metrics_JSON_SERIALISE-8         	   30000	     49681 ns/op	6835874.04 MB/s	   12436 B/op	       4 allocs/op
BenchmarkIt/127_metrics_PROTOBUF_PARSE-8         	  100000	     20086 ns/op	 301.20 MB/s	   14567 B/op	     267 allocs/op
BenchmarkIt/127_metrics_JSON_PARSE-8             	   10000	    195117 ns/op	  58.02 MB/s	   31636 B/op	     524 allocs/op
BenchmarkIt/127_metrics_JSON_VALIDATE-PARSE-8    	    2000	   1013845 ns/op	  11.17 MB/s	  250654 B/op	    8036 allocs/op
BenchmarkIt/8kb_msg_PROTOBUF_SERIALISE-8         	  200000	      6363 ns/op	260115970.45 MB/s	    9471 B/op	       1 allocs/op
BenchmarkIt/8kb_msg_JSON_SERIALISE-8             	  100000	     23792 ns/op	42037364.78 MB/s	   10387 B/op	       4 allocs/op
BenchmarkIt/8kb_msg_PROTOBUF_PARSE-8             	  200000	      6611 ns/op	1251.72 MB/s	    8527 B/op	       7 allocs/op
BenchmarkIt/8kb_msg_JSON_PARSE-8                 	   10000	    103502 ns/op	  96.63 MB/s	   19022 B/op	      14 allocs/op
BenchmarkIt/8kb_msg_JSON_VALIDATE-PARSE-8        	   10000	    225387 ns/op	  44.38 MB/s	   41629 B/op	     146 allocs/op
BenchmarkIt/3_fields_PROTOBUF_SERIALISE-8        	 2000000	       960 ns/op	1218617057.70 MB/s	     639 B/op	       1 allocs/op
BenchmarkIt/3_fields_JSON_SERIALISE-8            	  500000	      3062 ns/op	146294660.75 MB/s	    1040 B/op	       4 allocs/op
BenchmarkIt/3_fields_PROTOBUF_PARSE-8            	 1000000	      1383 ns/op	 422.81 MB/s	    1063 B/op	      17 allocs/op
BenchmarkIt/3_fields_JSON_PARSE-8                	  100000	     12202 ns/op	  73.43 MB/s	    1903 B/op	      26 allocs/op
BenchmarkIt/3_fields_JSON_VALIDATE-PARSE-8       	   30000	     47358 ns/op	  18.92 MB/s	    9722 B/op	     261 allocs/op
BenchmarkIt/255_fields_PROTOBUF_SERIALISE-8      	   30000	     50718 ns/op	25386829.48 MB/s	   49150 B/op	       1 allocs/op
BenchmarkIt/255_fields_JSON_SERIALISE-8          	   10000	    161095 ns/op	3368351.33 MB/s	   57516 B/op	       4 allocs/op
BenchmarkIt/255_fields_PROTOBUF_PARSE-8          	   20000	     71111 ns/op	 603.57 MB/s	   61540 B/op	     779 allocs/op
BenchmarkIt/255_fields_JSON_PARSE-8              	    2000	    705257 ns/op	  76.95 MB/s	  126368 B/op	    1009 allocs/op
BenchmarkIt/255_fields_JSON_VALIDATE-PARSE-8     	    1000	   2034334 ns/op	  26.68 MB/s	  444303 B/op	    9917 allocs/op
BenchmarkIt/255_metrics_PROTOBUF_SERIALISE-8     	   50000	     26322 ns/op	22919177.19 MB/s	   12287 B/op	       1 allocs/op
BenchmarkIt/255_metrics_JSON_SERIALISE-8         	   20000	     99347 ns/op	4480805.22 MB/s	   24730 B/op	       4 allocs/op
BenchmarkIt/255_metrics_PROTOBUF_PARSE-8         	   50000	     39286 ns/op	 307.13 MB/s	   28903 B/op	     524 allocs/op
BenchmarkIt/255_metrics_JSON_PARSE-8             	    5000	    379153 ns/op	  58.71 MB/s	   66578 B/op	    1009 allocs/op
BenchmarkIt/255_metrics_JSON_VALIDATE-PARSE-8    	    1000	   1985382 ns/op	  11.21 MB/s	  506555 B/op	   16032 allocs/op
BenchmarkIt/no_msg_PROTOBUF_SERIALISE-8          	 5000000	       307 ns/op	1318778050.88 MB/s	      95 B/op	       0 allocs/op
BenchmarkIt/no_msg_JSON_SERIALISE-8              	 1000000	      1553 ns/op	158978695.15 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/no_msg_PROTOBUF_PARSE-8              	 5000000	       392 ns/op	 206.32 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/no_msg_JSON_PARSE-8                  	  300000	      4776 ns/op	  51.71 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/no_msg_JSON_VALIDATE-PARSE-8         	   50000	     24262 ns/op	  10.18 MB/s	    4722 B/op	     141 allocs/op
BenchmarkIt/7_metrics_PROTOBUF_SERIALISE-8       	 1000000	      1074 ns/op	381706389.43 MB/s	     415 B/op	       0 allocs/op
BenchmarkIt/7_metrics_JSON_SERIALISE-8           	  300000	      4352 ns/op	60521027.96 MB/s	    1040 B/op	       4 allocs/op
BenchmarkIt/7_metrics_PROTOBUF_PARSE-8           	 1000000	      1714 ns/op	 239.12 MB/s	    1127 B/op	      23 allocs/op
BenchmarkIt/7_metrics_JSON_PARSE-8               	  100000	     15897 ns/op	  55.23 MB/s	    2231 B/op	      46 allocs/op
BenchmarkIt/7_metrics_JSON_VALIDATE-PARSE-8      	   20000	     82096 ns/op	  10.69 MB/s	   18129 B/op	     586 allocs/op
BenchmarkIt/31_metrics_PROTOBUF_SERIALISE-8      	  500000	      3585 ns/op	214492012.74 MB/s	    1791 B/op	       0 allocs/op
BenchmarkIt/31_metrics_JSON_SERIALISE-8          	  100000	     13539 ns/op	21396888.70 MB/s	    3217 B/op	       4 allocs/op
BenchmarkIt/31_metrics_PROTOBUF_PARSE-8          	  300000	      5501 ns/op	 279.55 MB/s	    3815 B/op	      73 allocs/op
BenchmarkIt/31_metrics_JSON_PARSE-8              	   30000	     50969 ns/op	  56.84 MB/s	    8871 B/op	     145 allocs/op
BenchmarkIt/31_metrics_JSON_VALIDATE-PARSE-8     	    5000	    268097 ns/op	  10.81 MB/s	   64899 B/op	    2073 allocs/op
BenchmarkIt/511_metrics_PROTOBUF_SERIALISE-8     	   30000	     52631 ns/op	13735435.48 MB/s	   24575 B/op	       0 allocs/op
BenchmarkIt/511_metrics_JSON_SERIALISE-8         	   10000	    200422 ns/op	2230865.17 MB/s	   49322 B/op	       4 allocs/op
BenchmarkIt/511_metrics_PROTOBUF_PARSE-8         	   20000	     78111 ns/op	 308.51 MB/s	   57573 B/op	    1037 allocs/op
BenchmarkIt/511_metrics_JSON_PARSE-8             	    2000	    781475 ns/op	  57.22 MB/s	  140658 B/op	    2016 allocs/op
BenchmarkIt/511_metrics_JSON_VALIDATE-PARSE-8    	     500	   3994003 ns/op	  11.20 MB/s	 1024311 B/op	   32077 allocs/op
BenchmarkIt/0_fields_PROTOBUF_SERIALISE-8        	 5000000	       308 ns/op	1310984129.99 MB/s	      95 B/op	       0 allocs/op
BenchmarkIt/0_fields_JSON_SERIALISE-8            	 1000000	      1546 ns/op	160395179.97 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/0_fields_PROTOBUF_PARSE-8            	 5000000	       390 ns/op	 207.36 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/0_fields_JSON_PARSE-8                	  300000	      4715 ns/op	  52.60 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/0_fields_JSON_VALIDATE-PARSE-8       	  100000	     23421 ns/op	  10.59 MB/s	    4723 B/op	     141 allocs/op
BenchmarkIt/15_fields_PROTOBUF_SERIALISE-8       	  500000	      3428 ns/op	379286151.63 MB/s	    2687 B/op	       0 allocs/op
BenchmarkIt/15_fields_JSON_SERIALISE-8           	  200000	      8591 ns/op	78425048.82 MB/s	    3601 B/op	       4 allocs/op
BenchmarkIt/15_fields_PROTOBUF_PARSE-8           	  300000	      4862 ns/op	 534.95 MB/s	    3943 B/op	      55 allocs/op
BenchmarkIt/15_fields_JSON_PARSE-8               	   30000	     44082 ns/op	  76.43 MB/s	    7871 B/op	      79 allocs/op
BenchmarkIt/15_fields_JSON_VALIDATE-PARSE-8      	   10000	    143931 ns/op	  23.41 MB/s	   29975 B/op	     721 allocs/op
BenchmarkIt/0_metrics_PROTOBUF_SERIALISE-8       	 5000000	       304 ns/op	1329216660.12 MB/s	      95 B/op	       0 allocs/op
BenchmarkIt/0_metrics_JSON_SERIALISE-8           	 1000000	      1544 ns/op	160616222.27 MB/s	     400 B/op	       4 allocs/op
BenchmarkIt/0_metrics_PROTOBUF_PARSE-8           	 5000000	       390 ns/op	 207.68 MB/s	     335 B/op	       5 allocs/op
BenchmarkIt/0_metrics_JSON_PARSE-8               	  300000	      4699 ns/op	  52.78 MB/s	     591 B/op	      12 allocs/op
BenchmarkIt/0_metrics_JSON_VALIDATE-PARSE-8      	  100000	     23534 ns/op	  10.54 MB/s	    4722 B/op	     141 allocs/op
BenchmarkIt/15_metrics_PROTOBUF_SERIALISE-8      	 1000000	      1903 ns/op	412915175.24 MB/s	     895 B/op	       0 allocs/op
BenchmarkIt/15_metrics_JSON_SERIALISE-8          	  200000	      7337 ns/op	42440835.81 MB/s	    1936 B/op	       4 allocs/op
BenchmarkIt/15_metrics_PROTOBUF_PARSE-8          	  500000	      3047 ns/op	 257.92 MB/s	    2023 B/op	      40 allocs/op
BenchmarkIt/15_metrics_JSON_PARSE-8              	   50000	     27394 ns/op	  56.84 MB/s	    4447 B/op	      81 allocs/op
BenchmarkIt/15_metrics_JSON_VALIDATE-PARSE-8     	   10000	    144288 ns/op	  10.79 MB/s	   33759 B/op	    1085 allocs/op
BenchmarkIt/63_metrics_PROTOBUF_SERIALISE-8      	  200000	      6746 ns/op	90181317.72 MB/s	    3071 B/op	       1 allocs/op
BenchmarkIt/63_metrics_JSON_SERIALISE-8          	   50000	     24842 ns/op	11451708.79 MB/s	    6290 B/op	       4 allocs/op
BenchmarkIt/63_metrics_PROTOBUF_PARSE-8          	  200000	     10343 ns/op	 294.10 MB/s	    7399 B/op	     138 allocs/op
BenchmarkIt/63_metrics_JSON_PARSE-8              	   20000	     98211 ns/op	  57.94 MB/s	   14655 B/op	     269 allocs/op
BenchmarkIt/63_metrics_JSON_VALIDATE-PARSE-8     	    3000	    513531 ns/op	  11.08 MB/s	  124347 B/op	    4048 allocs/op
BenchmarkIt/small_msg_PROTOBUF_SERIALISE-8       	 5000000	       332 ns/op	1593212132.66 MB/s	     111 B/op	       1 allocs/op
BenchmarkIt/small_msg_JSON_SERIALISE-8           	 1000000	      1639 ns/op	165285018.04 MB/s	     432 B/op	       4 allocs/op
BenchmarkIt/small_msg_PROTOBUF_PARSE-8           	 3000000	       445 ns/op	 238.07 MB/s	     367 B/op	       6 allocs/op
BenchmarkIt/small_msg_JSON_PARSE-8               	  300000	      4976 ns/op	  54.45 MB/s	     623 B/op	      13 allocs/op
BenchmarkIt/small_msg_JSON_VALIDATE-PARSE-8      	  100000	     24058 ns/op	  11.26 MB/s	    4803 B/op	     144 allocs/op
BenchmarkIt/63_fields_PROTOBUF_SERIALISE-8       	  100000	     12872 ns/op	82847399.02 MB/s	   10879 B/op	       0 allocs/op
BenchmarkIt/63_fields_JSON_SERIALISE-8           	   50000	     38157 ns/op	17845401.85 MB/s	   14485 B/op	       4 allocs/op
BenchmarkIt/63_fields_PROTOBUF_PARSE-8           	  100000	     18162 ns/op	 587.19 MB/s	   15463 B/op	     201 allocs/op
BenchmarkIt/63_fields_JSON_PARSE-8               	   10000	    177657 ns/op	  76.66 MB/s	   30364 B/op	     272 allocs/op
BenchmarkIt/63_fields_JSON_VALIDATE-PARSE-8      	    3000	    528902 ns/op	  25.75 MB/s	  110868 B/op	    2542 allocs/op
BenchmarkIt/127_fields_PROTOBUF_SERIALISE-8      	   50000	     25199 ns/op	42493219.09 MB/s	   21759 B/op	       0 allocs/op
BenchmarkIt/127_fields_JSON_SERIALISE-8          	   20000	     78984 ns/op	6868894.16 MB/s	   27419 B/op	       4 allocs/op
BenchmarkIt/127_fields_PROTOBUF_PARSE-8          	   50000	     35911 ns/op	 596.38 MB/s	   30823 B/op	     394 allocs/op
BenchmarkIt/127_fields_JSON_PARSE-8              	    5000	    353148 ns/op	  76.82 MB/s	   62179 B/op	     524 allocs/op
BenchmarkIt/127_fields_JSON_VALIDATE-PARSE-8     	    2000	   1044917 ns/op	  25.96 MB/s	  221221 B/op	    4989 allocs/op
BenchmarkIt/511_fields_PROTOBUF_SERIALISE-8      	   20000	    100598 ns/op	17082689.13 MB/s	   90107 B/op	       1 allocs/op
BenchmarkIt/511_fields_JSON_SERIALISE-8          	    5000	    323868 ns/op	1674063.92 MB/s	  114918 B/op	       4 allocs/op
BenchmarkIt/511_fields_PROTOBUF_PARSE-8          	   10000	    141903 ns/op	 605.55 MB/s	  122971 B/op	    1548 allocs/op
BenchmarkIt/511_fields_JSON_PARSE-8              	    1000	   1411449 ns/op	  76.84 MB/s	  263384 B/op	    2016 allocs/op
BenchmarkIt/511_fields_JSON_VALIDATE-PARSE-8     	     300	   4029391 ns/op	  26.92 MB/s	  899289 B/op	   19813 allocs/op
BenchmarkIt/1_metrics_PROTOBUF_SERIALISE-8       	 3000000	       439 ns/op	872922857.77 MB/s	     128 B/op	       1 allocs/op
BenchmarkIt/1_metrics_JSON_SERIALISE-8           	 1000000	      1998 ns/op	169091825.15 MB/s	     496 B/op	       4 allocs/op
BenchmarkIt/1_metrics_PROTOBUF_PARSE-8           	 3000000	       569 ns/op	 224.73 MB/s	     439 B/op	       8 allocs/op
BenchmarkIt/1_metrics_JSON_PARSE-8               	  200000	      6397 ns/op	  52.83 MB/s	     903 B/op	      18 allocs/op
BenchmarkIt/1_metrics_JSON_VALIDATE-PARSE-8      	   50000	     33082 ns/op	  10.22 MB/s	    6743 B/op	     207 allocs/op
BenchmarkIt/2kb_msg_PROTOBUF_SERIALISE-8         	 1000000	      1874 ns/op	1137137479.78 MB/s	    2303 B/op	       1 allocs/op
BenchmarkIt/2kb_msg_JSON_SERIALISE-8             	  200000	      6741 ns/op	80956698.99 MB/s	    3217 B/op	       4 allocs/op
BenchmarkIt/2kb_msg_PROTOBUF_PARSE-8             	 1000000	      2002 ns/op	1064.65 MB/s	    2383 B/op	       7 allocs/op
BenchmarkIt/2kb_msg_JSON_PARSE-8                 	   50000	     28853 ns/op	  94.58 MB/s	    5327 B/op	      14 allocs/op
BenchmarkIt/2kb_msg_JSON_VALIDATE-PARSE-8        	   20000	     73623 ns/op	  37.07 MB/s	   14219 B/op	     146 allocs/op
PASS
ok  	github.com/Scout24/protobuf-vs-json	256.535s
```
