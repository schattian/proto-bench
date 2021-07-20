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
Benchmark_GogoV1_JSON_Marshal-12         |      14196 |  84707 ns/op |   843 | 29660 |   1.20 |    1196 |    2.86
Benchmark_GogoV1_JSON_Unmarshal-12       |      10000 | 110846 ns/op |   843 | 27984 |   1.11 |     843 |    3.96
Benchmark_GogoV1_Proto_Marshal-12        |    1000000 |   1384 ns/op |   300 | 319 |   1.38 |   30010 |    4.34
Benchmark_GogoV1_Proto_Unmarshal-12      |    1000000 |   1428 ns/op |   300 | 709 |   1.43 |   30010 |    2.01
Benchmark_GoV1_JSON_Marshal-12           |      50515 |  24018 ns/op |   841 | 4801 |   1.21 |    4249 |    5.00
Benchmark_GoV1_JSON_Unmarshal-12         |      12873 |  89467 ns/op |   841 | 25039 |   1.15 |    1082 |    3.57
Benchmark_GoV1old_JSON_Marshal-12        |      15993 |  74710 ns/op |   846 | 24104 |   1.19 |    1353 |    3.10
Benchmark_GoV1old_JSON_Unmarshal-12      |      10000 | 109489 ns/op |   846 | 28318 |   1.09 |     846 |    3.87
Benchmark_GoV1old_Proto_Marshal-12       |     603187 |   2814 ns/op |   305 | 327 |   1.70 |   18403 |    8.61
Benchmark_GoV1old_Proto_Unmarshal-12     |     623715 |   2265 ns/op |   305 | 944 |   1.41 |   19029 |    2.40
Benchmark_GoV1_Proto_Marshal-12          |     617084 |   2654 ns/op |   300 | 319 |   1.64 |   18518 |    8.32
Benchmark_GoV1_Proto_Unmarshal-12        |     561639 |   2486 ns/op |   300 | 933 |   1.40 |   16854 |    2.66
Benchmark_GoV2_JSON_Marshal-12           |      73975 |  16437 ns/op |   841 | 3592 |   1.22 |    6223 |    4.58
Benchmark_GoV2_JSON_Unmarshal-12         |      45798 |  26259 ns/op |   841 | 3095 |   1.20 |    3852 |    8.48
Benchmark_GoV2_Proto_Marshal-12          |     618963 |   2597 ns/op |   300 | 319 |   1.61 |   18575 |    8.14
Benchmark_GoV2_Proto_Unmarshal-12        |     527601 |   2663 ns/op |   300 | 933 |   1.41 |   15833 |    2.85
Benchmark_GoV2_Proto_VTProto_Marshal-12  |    1000000 |   1450 ns/op |   300 | 319 |   1.45 |   30010 |    4.55
Benchmark_GoV2_Proto_VTProto_Unmarshal-12 |   10147508 |    153 ns/op |   300 |   0 |   1.56 |  304526 |    0.00


Totals:


benchmark                                | iter  | time/iter | bytes/op  |  allocs/op |tt.sec  | tt.kb        | ns/alloc
-----------------------------------------|-------|-----------|-----------|------------|--------|--------------|-----------
Benchmark_GoV2_Proto_VTProto_-12         |   11147508 |   1603 ns/op |   600 | 319 |  17.88 |  669073 |    5.03
Benchmark_GogoV1_Proto_-12               |    2000000 |   2812 ns/op |   600 | 1028 |   5.62 |  120040 |    2.74
Benchmark_GoV1old_Proto_-12              |    1226902 |   5079 ns/op |   610 | 1271 |   6.23 |   74865 |    4.00
Benchmark_GoV1_Proto_-12                 |    1178723 |   5140 ns/op |   600 | 1252 |   6.06 |   70746 |    4.11
Benchmark_GoV2_Proto_-12                 |    1146564 |   5260 ns/op |   600 | 1252 |   6.03 |   68816 |    4.20
Benchmark_GoV2_JSON_-12                  |     119773 |  42696 ns/op |  1682 | 6687 |   5.11 |   20153 |    6.38
Benchmark_GoV1_JSON_-12                  |      63388 | 113485 ns/op |  1682 | 29840 |   7.19 |   10665 |    3.80
Benchmark_GoV1old_JSON_-12               |      25993 | 184199 ns/op |  1692 | 52422 |   4.79 |    4399 |    3.51
Benchmark_GogoV1_JSON_-12                |      24196 | 195553 ns/op |  1686 | 57644 |   4.73 |    4080 |    3.39
