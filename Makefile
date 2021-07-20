.PHONY: test all old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go

test:
	rm results.txt || true
	go test -count=1 -timeout 5m -bench=. | grep Benchmark_Go > results.txt
	cd old && go test -timeout 5m -count=1 -bench=. | grep Benchmark_Go >> ../results.txt
	sort -o results.txt results.txt
	cat results.txt

all: old/structdef-go-v1.pb.go structdef-gogo-v1.pb.go structdef-go-v1.pb.go structdef-go-v2.pb.go structdef-go-v2-vt-proto.pb.go

old/structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.5
	protoc --go_out=paths=source_relative:./old structdef-go-v1.proto

structdef-gogo-v1.pb.go: structdef-gogo-v1.proto
	go get github.com/gogo/protobuf/protoc-gen-gogofaster
	protoc --gogofaster_out=paths=source_relative:. structdef-gogo-v1.proto

structdef-go-v1.pb.go: structdef-go-v1.proto
	go get github.com/golang/protobuf/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. structdef-go-v1.proto

structdef-go-v2.pb.go: structdef-go-v2.proto
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:. structdef-go-v2.proto

structdef-go-v2-vt-proto.pb.go: structdef-go-v2.proto
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=paths=source_relative:.\
		--go-vtproto_out=$(GOPATH)/src/ \
		--plugin protoc-gen-go-vtproto="${GOBIN}/protoc-gen-go-vtproto" \
    	--go-vtproto_opt=features=marshal+unmarshal+size \
		structdef-go-v2.proto
	mv $(GOPATH)/src/github.com/alexshtin/protobench/structdef-go-v2_vtproto.pb.go .


clean:
	rm -f *.pb.go
	rm -f old/*.pb.go
