.PHONY: all test clean

test:
	rm results.txt || true
	go test -count=1 -bench=.  | grep Benchmark > results.txt
	sort -o results.txt results.txt
	cat results.txt

all: official gogo

official: def.proto
	docker run --rm -v $(CURDIR):/src:rw $(PROTOC_DOCKER_IMAGE) -I /src --go_out=paths=source_relative:./official def.proto

gogo: def.proto
	docker run --rm -v $(CURDIR):/src:rw $(PROTOC_DOCKER_IMAGE) -I=. --gogofaster_out=paths=source_relative:./gogo def.proto
 
clean:
	rm -rf gogo official 
