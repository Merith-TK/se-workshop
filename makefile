
build: 
	go build -o bin/sew.exe ./cmd/sew
	go build -o bin/se-chopshop.exe ./cmd/se-chopshop
	go build -o bin/sew-prefab.exe ./cmd/sew-prefab

build-sew:
	go build -o bin/sew.exe ./cmd/sew

build-chopshop:
	go build -o bin/se-chopshop.exe ./cmd/se-chopshop

build-prefab:
	go build -o bin/sew-prefab.exe ./cmd/sew-prefab

install:
	go install ./cmd/sew
	go install ./cmd/se-chopshop
	go install ./cmd/sew-prefab

install-sew:
	go install ./cmd/sew

install-chopshop:
	go install ./cmd/se-chopshop

install-prefab:
	go install ./cmd/sew-prefab

run-sew: build-sew
	./bin/sew.exe

run-chopshop: build-chopshop
	./bin/se-chopshop.exe

run-prefab: build-prefab
	./bin/sew-prefab.exe

test:
	go test ./...

clean:
	rm -rf bin/

.PHONY: build build-sew build-chopshop build-prefab install install-sew install-chopshop install-prefab run-sew run-chopshop run-prefab test clean