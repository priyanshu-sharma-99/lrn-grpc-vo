BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
SERVER_BIN = ${SERVER_DIR}
CLIENT_BIN = ${CLIENT_DIR}
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r

.PHONY: all greet calculator blog about

project := greet calculator blog

all: $(project)

greet calculator blog:
	@${CHECK_DIR_CMD}
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
#	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
#	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

greet_server_build:
	go build -o ./bin/greet/server/ ./greet/server/*

greet_client_build:
	go build -o ./bin/greet/client/ ./greet/client/*

greet_server_run:
	go run ./bin/greet/server/*

greet_client_run:
	go run ./bin/greet/client/*

calculator_server_build:
	go build -o ./bin/calculator/server/ ./calculator/server

calculator_client_build:
	go build -o ./bin/calculator/client/ ./calculator/client

calculator_server_run:
	go run ./calculator/server/*

calculator_client_run:
	go run ./calculator/client/*





about:
	@echo "Shell: ${SHELL}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"