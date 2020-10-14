all: clean build run

build:
	mkdir -p build/
	CGO_ENABLED=0 go build -ldflags '-w -s -extldflags "-static"' -o build/isa ./cmd

run:
	./build/isa --loglevel=debug --conf=test/isa.yaml

clean:
	rm -r ./build || true

update:
	go get -u all && go mod vendor && go mod tidy && go mod vendor
