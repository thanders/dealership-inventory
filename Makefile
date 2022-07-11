BIN = dealership-inventory
PROTO_DIR = proto

REPOSITORY = github.com/thanders/dealership-inventory

ifeq ($(OS), Windows_NT)
	OS = windows
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	PACKAGE = $(shell Get-Content go.mod -head 1 | Foreach-Object { $$data = $$_ -split " "; "{0}" -f $$data[1]})
else
	UNAME := $(shell uname -s)
	ifeq ($(UNAME),Darwin)
		OS = macos
	else ifeq ($(UNAME),Linux)
		OS = linux
	else
    	$(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

build: 	generate
	go build -o ${BIN} .

buildServer:
	go build -o greetServer ./greet/server

buildClient:
	go build -o greetClient ./greet/client

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

generateServer:
	protoc -Igreet/proto --go_out=. --go_opt=module=${REPOSITORY} --go-grpc_out=. --go-grpc_opt=module=${REPOSITORY} greet/proto/*.proto

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm ${BIN}