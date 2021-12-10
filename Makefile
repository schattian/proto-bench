test:
	rm results.txt || true
	go test -count=1 -bench=. | grep Benchmark_Go > results.txt
	cd old && go test -count=1 -bench=. | grep Benchmark_Go >> ../results.txt
	sort -o results.txt results.txt
	cat results.txt

all: official gogo

official: official.proto
	docker run --rm -v $(CURDIR):/src:rw $(PROTOC_DOCKER_IMAGE) -I /src --go_out=paths=source_relative:./official official.proto

gogo: gogo.proto
	docker run --rm -v $(CURDIR):/src:rw $(PROTOC_DOCKER_IMAGE) -I /src --gogofaster_out=paths=source_relative:./gogo gogo.proto

clean:
	rm -f *.pb.go
