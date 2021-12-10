test:
	rm results.txt || true
	go test -count=1 -bench=. | grep Benchmark_Go > results.txt
	cd old && go test -count=1 -bench=. | grep Benchmark_Go >> ../results.txt
	sort -o results.txt results.txt
	cat results.txt

all: official.pb.go gogo.pb.go

official.pb.go: official.proto
	docker run --rm -v $(pwd):/src:rw protoc -I /src --go_out=paths=source_relative:. official.proto

gogo.pb.go: gogo.proto
	docker run --rm -v $(pwd):/src:rw protoc -I /src --gogofaster_out=paths=source_relative:. gogo.proto

clean:
	rm -f *.pb.go
