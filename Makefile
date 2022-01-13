GOLANG_DOCKER_IMAGE := golang:1.17.6
APP_NAME := golang-tasks
BUILD_TAG ?= latest

build:
 GOARCH=amd64 GOOS=linux go build -o dist/${APP_NAME}:${BUILD_TAG} main.go

vet:	## vet code
	@echo "Vetting..."
	docker run --rm \
		-v "$$PWD":/usr/src/${APP_NAME} \
		-w /usr/src/${APP_NAME} \
		golang:1.17.6 \
		go vet ./...

fmt:	## format code
	@echo "Formatting..."
	docker run --rm \
		-v "$$PWD":/usr/src/${APP_NAME} \
		-w /usr/src/${APP_NAME} \
		golang:1.17.6 \
		go fmt ./...

test: vet fmt	## test
	@echo "Testing..."
	docker run --rm \
		-v "$$PWD":/usr/src/${APP_NAME} \
		-w /usr/src/${APP_NAME} \
		golang:1.17.6 \
		go test ./...

build:	## build binary
	@echo "Building binary..."
	docker run --rm \
    		-v "$$PWD":/usr/src/${APP_NAME} \
    		-w /usr/src/${APP_NAME} \
    		--env CGO_ENABLED=0 --env GOOS=linux \
    		golang:1.17.6 \
    		go build -o dist/${APP_NAME}:${BUILD_TAG}

stop:	## stop binary
	@-pkill ${APP_NAME}

run: stop build	## run binary
	@echo "Starting binary..."
	./dist/${APP_NAME}:${BUILD_TAG} &

build-docker-image:	## build docker image
	@echo "Building docker image..."
	docker build --pull -t ${APP_NAME}:${BUILD_TAG} .

run-docker: stop-docker build-docker-image ## run docker image
	@echo "Starting docker image..."
	docker run -d -p 8081:8081 ${APP_NAME}:${BUILD_TAG}

stop-docker:	## stop docker image
	@echo "Stopping existing docker image..."
	@-docker stop $(APP_NAME):${BUILD_TAG}; docker rm $(APP_NAME):${BUILD_TAG}



help: ## Help target
		@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help
.DEFAULT_GOAL := help
