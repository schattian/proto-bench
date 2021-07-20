# Benchmarks of Go Protobuf libraries

This is a test suite for benchmarking various Go Protobuf libraries.

## Libraries

* [Go Protobuf](https://blog.golang.org/protobuf-apiv2):
  - [github.com/golang/protobuf](https://github.com/golang/protobuf/releases/tag/v1.3.5) ~v1.3.5 is the most recent pre-APIv2 version of APIv1 (`old` sub-dir).
  - [github.com/golang/protobuf](https://github.com/golang/protobuf) ^v1.4.0 is a version of APIv1 implemented in terms of APIv2.
  - [google.golang.org/protobuf](https://github.com/protocolbuffers/protobuf-go) ^v1.20.0 is APIv2.

* [Gogo Protobuf](https://github.com/gogo/protobuf):
  - [github.com/gogo/protobuf](https://github.com/gogo/protobuf) Protocol Buffers for Go with Gadgets.

## Running the benchmarks

```bash
make all
make test
```

To update the table in the README:

```bash
./stats.sh
```

## Data

The data being serialized is the following structure with randomly generated values:

```proto
message Go {
  string name = 1;
  int64 birthDay = 2;
  string phone = 3;
  int32 siblings = 4;
  bool spouse = 5;
  double money = 6;
  Type type = 7;
  oneof values {
    string value_s = 8;
    int32 value_i = 9;
    double value_d = 10;
  }
}

enum Type {
  TYPE_UNSPECIFIED = 0;
  TYPE_R = 1;
  TYPE_S = 2;
}
```


## Results

benchmark                                | iter      | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-----------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GogoV1_JSON_Marshal-12         |      10000 | 106281 ns/op |  1063 | 37477 |   1.06 |    1063 |    2.84
Benchmark_GogoV1_JSON_Unmarshal-12       |       8308 | 140990 ns/op |  1063 | 35374 |   1.17 |     883 |    3.99
Benchmark_GogoV1_Proto_Marshal-12        |    1000000 |   1637 ns/op |   383 | 398 |   1.64 |   38370 |    4.11
Benchmark_GogoV1_Proto_Unmarshal-12      |     744663 |   1742 ns/op |   383 | 928 |   1.30 |   28572 |    1.88
Benchmark_GoV1_JSON_Marshal-12           |      39850 |  30378 ns/op |  1061 | 6736 |   1.21 |    4228 |    4.51
Benchmark_GoV1_JSON_Unmarshal-12         |      10000 | 114152 ns/op |  1061 | 31620 |   1.14 |    1061 |    3.61
Benchmark_GoV1old_JSON_Marshal-12        |      12524 |  93740 ns/op |  1066 | 30074 |   1.17 |    1335 |    3.12
Benchmark_GoV1old_JSON_Unmarshal-12      |       8535 | 138847 ns/op |  1066 | 35723 |   1.19 |     909 |    3.89
Benchmark_GoV1old_Proto_Marshal-12       |     471633 |   3376 ns/op |   388 | 417 |   1.59 |   18332 |    8.10
Benchmark_GoV1old_Proto_Unmarshal-12     |     482198 |   2751 ns/op |   388 | 1160 |   1.33 |   18743 |    2.37
Benchmark_GoV1_Proto_Marshal-12          |     479772 |   3164 ns/op |   383 | 398 |   1.52 |   18408 |    7.95
Benchmark_GoV1_Proto_Unmarshal-12        |     444213 |   2762 ns/op |   383 | 1152 |   1.23 |   17044 |    2.40
Benchmark_GoV2_JSON_Marshal-12           |      56851 |  21377 ns/op |  1119 | 4999 |   1.22 |    6361 |    4.28
Benchmark_GoV2_JSON_Unmarshal-12         |      36128 |  33217 ns/op |  1118 | 3865 |   1.20 |    4039 |    8.59
Benchmark_GoV2_Proto_Marshal-12          |     521704 |   3170 ns/op |   383 | 398 |   1.65 |   20017 |    7.96
Benchmark_GoV2_Proto_Unmarshal-12        |     444256 |   2823 ns/op |   383 | 1152 |   1.25 |   17046 |    2.45
Benchmark_GoV2_Proto_VTProto_Marshal-12  |    1000000 |   1843 ns/op |   383 | 398 |   1.84 |   38370 |    4.63
Benchmark_GoV2_Proto_VTProto_Unmarshal-12 |    8925570 |    151 ns/op |   383 |   0 |   1.35 |  342474 |    0.00


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GoV2_Proto_VTProto_-12         |    9925570 |   1994 ns/op |   767 | 398 |  19.80 |  761688 |    5.01
Benchmark_GogoV1_Proto_-12               |    1744663 |   3379 ns/op |   767 | 1326 |   5.90 |  133885 |    2.55
Benchmark_GoV1_Proto_-12                 |     923985 |   5926 ns/op |   767 | 1550 |   5.48 |   70906 |    3.82
Benchmark_GoV2_Proto_-12                 |     965960 |   5993 ns/op |   767 | 1550 |   5.79 |   74127 |    3.87
Benchmark_GoV1old_Proto_-12              |     953831 |   6127 ns/op |   777 | 1577 |   5.84 |   74150 |    3.89
Benchmark_GoV2_JSON_-12                  |      92979 |  54594 ns/op |  2237 | 8864 |   5.08 |   20799 |    6.16
Benchmark_GoV1_JSON_-12                  |      49850 | 144530 ns/op |  2122 | 38356 |   7.20 |   10578 |    3.77
Benchmark_GoV1old_JSON_-12               |      21059 | 232587 ns/op |  2132 | 65797 |   4.90 |    4489 |    3.53
Benchmark_GogoV1_JSON_-12                |      18308 | 247271 ns/op |  2126 | 72851 |   4.53 |    3892 |    3.39
