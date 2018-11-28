# Makefile

#export GOPATH := $(shell pwd)

all:
	echo $GOPATH
	#go get -d
	go build -o main

build:
	echo $GOPATH
	#go get -d
	go build -o out.bin
	
clean:
	go clean
