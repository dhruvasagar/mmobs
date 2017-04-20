BINDIR=bin

.PHONY: all clean build

all: build

bin:
	mkdir bin

build: bin
	go build -o $(BINDIR)/server src/server/*
	go build -o $(BINDIR)/client src/client/*

clean:
	@rm -f $(BINDIR)
