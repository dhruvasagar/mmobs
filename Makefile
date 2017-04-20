BINDIR=bin

.PHONY: all clean build server client

all: build

bin:
	@mkdir bin

build: bin
	go build -o $(BINDIR)/server src/server/*
	go build -o $(BINDIR)/client src/client/*

server: build
	go run src/server/*

client: build
	go run src/client/*

clean:
	@rm -rf $(BINDIR)
