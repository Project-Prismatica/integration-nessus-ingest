# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . prismatica_integration_nessus.proto --lile-server_out=. --go_out=plugins=grpc:${GOPATH}/src

test:
	go test -v ./... -cover

docker:
	docker build . -t prismatica_integration_nessus:latest
