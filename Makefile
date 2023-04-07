GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
GITHASH=$(shell git rev-parse HEAD 2>/dev/null)
BUILDAT=$(shell date +%FT%T%z)
API_PROTO_FILES=$(shell find api/* -name *.proto)
LDFLAGS="-s -w -X moredoc/cmd.GitHash=${GITHASH} -X moredoc/cmd.BuildAt=${BUILDAT} -X moredoc/cmd.Version=${VERSION}"

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest


.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--proto_path=./api \
		--gogofaster_out="plugins=grpc,paths=source_relative:." \
		--grpc-gateway_out="paths=source_relative:." \
		$(API_PROTO_FILES)

doc:
	for file in $(API_PROTO_FILES); do \
		protoc --proto_path=. \
		--proto_path=./third_party \
		--proto_path=./api \
		--doc_out=docs/api \
		--doc_opt=markdown,`basename $$file .proto`.md \
		--openapi_out==paths=source_relative:docs \
		$$file; \
	done
	# 整合到单文件
	protoc --proto_path=. \
		--proto_path=./third_party \
		--proto_path=./api \
		--doc_out=docs/api \
		--doc_opt=markdown,apis.md \
		--openapi_out==paths=source_relative:docs \
		$(API_PROTO_FILES)

.PHONY: clean-api-go
# clean api go file
clean-api-go:
	rm -rf api/*/*.go

builddarwin:
	GOOS=darwin GOARCH=amd64 go build -v -o release/${VERSION}/darwin/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${VERSION}/darwin
	cp -r dictionary release/${VERSION}/darwin
	cp -r app.example.toml release/${VERSION}/darwin
	rm -rf release/${VERSION}/darwin/dist/_nuxt/icons
	rm -rf release/${VERSION}/darwin/dist/_nuxt/manifest*

buildlinux:
	GOOS=linux GOARCH=amd64 go build -v -o release/${VERSION}/linux/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${VERSION}/linux
	cp -r dictionary release/${VERSION}/linux
	cp -r app.example.toml release/${VERSION}/linux
	rm -rf release/${VERSION}/linux/dist/_nuxt/icons
	rm -rf release/${VERSION}/linux/dist/_nuxt/manifest*

buildwin:
	GOOS=windows GOARCH=amd64 go build -v -o release/${VERSION}/windows/moredoc.exe -ldflags ${LDFLAGS}
	cp -r dist release/${VERSION}/windows
	cp -r dictionary release/${VERSION}/windows
	cp -r app.example.toml release/${VERSION}/windows
	rm -rf release/${VERSION}/windows/dist/_nuxt/icons
	rm -rf release/${VERSION}/windows/dist/_nuxt/manifest*

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
