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

buildServer:
	go build -o bin/server ./pkg/server

buildClient:
	go build -o bin/client ./pkg/client

generateServer:
	protoc -Ipkg/proto --go_out=. --go_opt=module=${REPOSITORY} --go-grpc_out=. --go-grpc_opt=module=${REPOSITORY} pkg/proto/*.proto

clean:
	-rm pkg/${PROTO_DIR}/*.pb.go
	-rm bin/client
	-rm bin/server
	-rm build/*.snap