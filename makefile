
build: 
	go build -o bin/main.exe .

run: build
	./bin/main.exe

steamcmd: build
	./bin/main.exe steamcmd