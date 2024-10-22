
build: 
	go build -o bin/sew.exe ./cmd/sew
install:
	go install ./cmd/sew
run: build
	./bin/main.exe

steamcmd: build
	./bin/main.exe steamcmd