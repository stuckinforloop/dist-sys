init:
	curl -L -O https://github.com/jepsen-io/maelstrom/releases/download/v0.2.4/maelstrom.tar.bz2
	tar -xjf maelstrom.tar.bz2
	rm maelstrom.tar.bz2
	go mod tidy
	go build -o node main.go

build:
	go build -o node main.go

run echo:
	./maelstrom/maelstrom test -w echo --bin ./node --node-count 1 --time-limit 10

debug:
	./maelstrom/maelstrom serve
