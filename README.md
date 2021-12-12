# Benchmarks of Go protobuf libs

This is a test suite for benchmarking various Go Protobuf libraries.

## Libraries

* [Official Go Protobuf](https://github.com/golang/protobuf):
  * [protojson (previously, jsonpb)](https://google.golang.org/protobuf/encoding/protojson) 

* [Gogo Protobuf](https://github.com/gogo/protobuf):
  * [gogojsonpb](https://github.com/gogo/protobuf/tree/master/jsonpb)

## Usage 

```bash
SIZE=n make all
SIZE=n make test
```

Where size is a modifier over the pb structs used in tests (default: 1)


 - To update the table in the README:

```bash
./stats.sh
```

 - To compile the proto def:

```bash
PROTOC_DOCKER_IMAGE=$YOUR_DOCKER_IMAGE make all 
``` 
