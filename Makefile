all: clean build run

build:
	mkdir -p build/
	go build -o build/isa ./cmd

run:
	./build/isa --loglevel=debug --conf=test/isa.yaml

clean:
	rm -r ./build || true
