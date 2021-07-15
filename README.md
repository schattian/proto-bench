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
Benchmark_GogoV1_JSON_Marshal-12         |      84378 |  13499 ns/op |   155 | 4504 |   1.14 |    1314 |    3.00
Benchmark_GogoV1_JSON_Unmarshal-12       |      73035 |  16106 ns/op |   155 | 4297 |   1.18 |    1138 |    3.75
Benchmark_GogoV1_Proto_Marshal-12        |    1956818 |    643 ns/op |    60 |  64 |   1.26 |   11760 |   10.05
Benchmark_GogoV1_Proto_Unmarshal-12      |    3228058 |    415 ns/op |    60 |  48 |   1.34 |   19400 |    8.65
Benchmark_GoV1_JSON_Marshal-12           |     275522 |   4549 ns/op |   153 | 972 |   1.25 |    4240 |    4.68
Benchmark_GoV1_JSON_Unmarshal-12         |      91419 |  12530 ns/op |   154 | 3799 |   1.15 |    1407 |    3.30
Benchmark_GoV1old_JSON_Marshal-12        |      95300 |  12447 ns/op |   154 | 3960 |   1.19 |    1467 |    3.14
Benchmark_GoV1old_JSON_Unmarshal-12      |      75519 |  16159 ns/op |   154 | 4376 |   1.22 |    1162 |    3.69
Benchmark_GoV1old_Proto_Marshal-12       |    1000000 |   1097 ns/op |    60 |  72 |   1.10 |    6010 |   15.24
Benchmark_GoV1old_Proto_Unmarshal-12     |    1749964 |    701 ns/op |    60 | 160 |   1.23 |   10517 |    4.38
Benchmark_GoV1_Proto_Marshal-12          |    1000000 |   1143 ns/op |    60 |  64 |   1.14 |    6010 |   17.86
Benchmark_GoV1_Proto_Unmarshal-12        |    1585590 |    744 ns/op |    60 | 176 |   1.18 |    9529 |    4.23
Benchmark_GoV2_JSON_Marshal-12           |     374341 |   3443 ns/op |   159 | 919 |   1.29 |    5985 |    3.75
Benchmark_GoV2_JSON_Unmarshal-12         |     265340 |   4495 ns/op |   159 | 690 |   1.19 |    4242 |    6.51
Benchmark_GoV2_Proto_Marshal-12          |    1000000 |   1059 ns/op |    60 |  64 |   1.06 |    6010 |   16.55
Benchmark_GoV2_Proto_Unmarshal-12        |    1414940 |    834 ns/op |    60 | 176 |   1.18 |    8503 |    4.74
Benchmark_GoV2_Proto_VTProto_Marshal-12  |    1935825 |    705 ns/op |    60 |  64 |   1.37 |   11634 |   11.02
Benchmark_GoV2_Proto_VTProto_Unmarshal-12 |   10005579 |    139 ns/op |    60 |   0 |   1.39 |   60133 |    0.00


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GoV2_Proto_VTProto_-12         |   11941404 |    844 ns/op |   120 |  64 |  10.09 |  143535 |   13.20
Benchmark_GogoV1_Proto_-12               |    5184876 |   1058 ns/op |   120 | 112 |   5.49 |   62322 |    9.45
Benchmark_GoV1old_Proto_-12              |    2749964 |   1798 ns/op |   120 | 232 |   4.95 |   33054 |    7.75
Benchmark_GoV1_Proto_-12                 |    2585590 |   1887 ns/op |   120 | 240 |   4.88 |   31078 |    7.86
Benchmark_GoV2_Proto_-12                 |    2414940 |   1893 ns/op |   120 | 240 |   4.57 |   29027 |    7.89
Benchmark_GoV2_JSON_-12                  |     639681 |   7938 ns/op |   319 | 1609 |   5.08 |   20456 |    4.93
Benchmark_GoV1_JSON_-12                  |     366941 |  17079 ns/op |   307 | 4771 |   6.27 |   11298 |    3.58
Benchmark_GoV1old_JSON_-12               |     170819 |  28606 ns/op |   308 | 8336 |   4.89 |    5261 |    3.43
Benchmark_GogoV1_JSON_-12                |     157413 |  29605 ns/op |   311 | 8801 |   4.66 |    4906 |    3.36
