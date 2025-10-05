
build: 
	go build -o bin/sew.exe ./cmd/sew
	go build -o bin/se-chopshop.exe ./cmd/se-chopshop

build-sew:
	go build -o bin/sew.exe ./cmd/sew

build-chopshop:
	go build -o bin/se-chopshop.exe ./cmd/se-chopshop

install:
	go install ./cmd/sew
	go install ./cmd/se-chopshop

install-sew:
	go install ./cmd/sew

install-chopshop:
	go install ./cmd/se-chopshop

run-sew: build-sew
	./bin/sew.exe

run-chopshop: build-chopshop
	./bin/se-chopshop.exe

test:
	go test ./...

clean:
	rm -rf bin/

.PHONY: build build-sew build-chopshop install install-sew install-chopshop run-sew run-chopshop test clean