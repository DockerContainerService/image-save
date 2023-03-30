NAME = imsave
VERSION = v1.0.0-release_build.0

OS = linux
architecture = amd64 arm64

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

build: deps ## Build the project
	go build -ldflags "-s -w -X 'github.com/DockerContainerService/image-save/cmd.version=$(VERSION)' -X 'main.logLevel=4'"

all: release release-windows ## Generate releases for all supported systems

release: clean deps ## Generate releases for unix systems
	@for arch in $(architecture);\
	do \
		for os in linux;\
		do \
			echo "Building $$os-$$arch"; \
			mkdir -p build; \
			GOOS=$$os GOARCH=$$arch go build -ldflags "-s -w -X 'github.com/DockerContainerService/image-save/cmd.version=$(VERSION)' -X 'main.logLevel=4'" -o build/$(NAME)-$$os-$$arch; \
			upx -9 build/$(NAME)-$$os-$$arch; \
		done \
	done

release-windows: clean deps ## Generate releases for unix systems
	@for arch in $(architecture);\
	do \
		for os in windows;\
		do \
			echo "Building $$os-$$arch"; \
			mkdir -p build; \
			GOOS=$$os GOARCH=$$arch go build -ldflags "-s -w -X 'github.com/DockerContainerService/image-save/cmd.version=$(VERSION)' -X 'main.logLevel=4'" -o build/$(NAME)-$$os-$$arch.exe; \
			upx -9 build/$(NAME)-$$os-$$arch.exe; \
		done \
	done

test: deps ## Execute tests
	go test ./...

deps: ## Install dependencies using go get
	go get -d -v -t ./...

clean: ## Remove building artifacts
	rm -rf build
	rm -f $(NAME)
