GOPATH:=$(shell go env GOPATH)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
VERSION=$(shell git tag --sort=-v:refname | grep ${BRANCH}- | head -n 1 | sed 's/${BRANCH}-//')
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

# 生成openapi
openapi:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--proto_path=./api \
		--openapiv2_out ./docs \
		$(API_PROTO_FILES)

.PHONY: clean-api-go
# clean api go file
clean-api-go:
	rm -rf api/*/*.go

builddarwin:
	rm -rf release/${BRANCH}/${VERSION}/darwin
	GOOS=darwin GOARCH=amd64 go build -v -o release/${BRANCH}/${VERSION}/darwin/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/darwin
	cp -r dictionary release/${BRANCH}/${VERSION}/darwin
	cp -r app.example.toml release/${BRANCH}/${VERSION}/darwin
	rm -rf release/${BRANCH}/${VERSION}/darwin/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/darwin/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/darwin/ && tar -zcvf ../moredoc_cics_${VERSION}_darwin_amd64.tar.gz ./* && cd ../../

builddarwinarm:
	rm -rf release/${BRANCH}/${VERSION}/darwin-arm
	GOOS=darwin GOARCH=arm64 go build -v -o release/${BRANCH}/${VERSION}/darwin-arm/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/darwin-arm
	cp -r dictionary release/${BRANCH}/${VERSION}/darwin-arm
	cp -r app.example.toml release/${BRANCH}/${VERSION}/darwin-arm
	rm -rf release/${BRANCH}/${VERSION}/darwin-arm/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/darwin-arm/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/darwin-arm/ && tar -zcvf ../moredoc_pro_${VERSION}_darwin_arm64.tar.gz ./* && cd ../../

buildlinux:
	rm -rf release/${BRANCH}/${VERSION}/linux
	GOOS=linux GOARCH=amd64 go build -v -o release/${BRANCH}/${VERSION}/linux/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/linux
	cp -r dictionary release/${BRANCH}/${VERSION}/linux
	cp -r app.example.toml release/${BRANCH}/${VERSION}/linux
	rm -rf release/${BRANCH}/${VERSION}/linux/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/linux/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/linux/ && tar -zcvf ../moredoc_cics_${VERSION}_linux_amd64.tar.gz ./* && cd ../../

builddockeramd:
	rm -rf release/${BRANCH}/${VERSION}/dockeramd
	GOOS=linux GOARCH=amd64 go build -v -o release/${BRANCH}/${VERSION}/dockeramd/server/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/dockeramd/server
	cp -r dictionary release/${BRANCH}/${VERSION}/dockeramd/server
	cp docker/dockerfile release/${BRANCH}/${VERSION}/dockeramd/dockerfile
	cp docker/docker-compose.yml release/${BRANCH}/${VERSION}/dockeramd/docker-compose.yml
	cp docker/README.md release/${BRANCH}/${VERSION}/dockeramd/部署教程.md
	cp -r docker/mysql release/${BRANCH}/${VERSION}/dockeramd/mysql
	rm -rf release/${BRANCH}/${VERSION}/dockeramd/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/dockeramd/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/dockeramd/ && tar -zcvf ../moredoc_pro_${VERSION}_docker_amd64.tar.gz ./* && cd ../../


builddockerarm:
	rm -rf release/${BRANCH}/${VERSION}/dockerarm
	GOOS=linux GOARCH=arm64 go build -v -o release/${BRANCH}/${VERSION}/dockerarm/server/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/dockerarm/server
	cp -r dictionary release/${BRANCH}/${VERSION}/dockerarm/server
	cp docker/dockerfile release/${BRANCH}/${VERSION}/dockerarm/dockerfile
	cp docker/docker-compose.yml release/${BRANCH}/${VERSION}/dockerarm/docker-compose.yml
	cp docker/README.md release/${BRANCH}/${VERSION}/dockerarm/部署教程.md
	cp -r docker/mysql release/${BRANCH}/${VERSION}/dockerarm/mysql
	rm -rf release/${BRANCH}/${VERSION}/dockerarm/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/dockerarm/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/dockerarm/ && tar -zcvf ../moredoc_pro_${VERSION}_docker_arm64.tar.gz ./* && cd ../../


buildwin:
	rm -rf release/${BRANCH}/${VERSION}/windows
	GOOS=windows GOARCH=amd64 go build -v -o release/${BRANCH}/${VERSION}/windows/moredoc.exe -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/windows
	cp -r dictionary release/${BRANCH}/${VERSION}/windows
	cp -r app.example.toml release/${BRANCH}/${VERSION}/windows
	rm -rf release/${BRANCH}/${VERSION}/windows/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/windows/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/windows/ && tar -zcvf ../moredoc_cics_${VERSION}_windows_amd64.tar.gz ./* && cd ../../

buildlinuxarm:
	rm -rf release/${BRANCH}/${VERSION}/linux-arm
	GOOS=linux GOARCH=arm64 go build -v -o release/${BRANCH}/${VERSION}/linux-arm/moredoc -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/linux-arm
	cp -r dictionary release/${BRANCH}/${VERSION}/linux-arm
	cp -r app.example.toml release/${BRANCH}/${VERSION}/linux-arm
	rm -rf release/${BRANCH}/${VERSION}/linux-arm/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/linux-arm/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/linux-arm/ && tar -zcvf ../moredoc_cics_${VERSION}_linux_arm64.tar.gz ./* && cd ../../

buildwinarm:
	rm -rf release/${BRANCH}/${VERSION}/windows-arm
	GOOS=windows GOARCH=arm64 go build -v -o release/${BRANCH}/${VERSION}/windows-arm/moredoc.exe -ldflags ${LDFLAGS}
	cp -r dist release/${BRANCH}/${VERSION}/windows-arm
	cp -r dictionary release/${BRANCH}/${VERSION}/windows-arm
	cp -r app.example.toml release/${BRANCH}/${VERSION}/windows-arm
	rm -rf release/${BRANCH}/${VERSION}/windows-arm/dist/_nuxt/icons
	rm -rf release/${BRANCH}/${VERSION}/windows-arm/dist/_nuxt/manifest*
	cd release/${BRANCH}/${VERSION}/windows-arm/ && tar -zcvf ../moredoc_cics_${VERSION}_windows_arm64.tar.gz ./* && cd ../../

# 一键编译所有平台
buildall: builddarwin builddarwinarm buildlinux buildwin buildlinuxarm buildwinarm builddockerarm builddockeramd

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
