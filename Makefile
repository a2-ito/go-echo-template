DOCKER		= docker
DOCKERFILE	= Dockerfile
APP		= itotest
BINARY		= go-echo-template
LINT_IGNORE	= ""

export
PORT	:= 8080

all: format lint run

lint:
	staticcheck ./...
	golangci-lint run

hadolint:
	$(DOCKER) run --rm -i hadolint/hadolint hadolint - --ignore ${LINT_IGNORE} < $(DOCKERFILE)

format:
	gofmt -l -w .

run:
	go run main.go

build:
	go build

build-container:
	pack build --builder=gcr.io/buildpacks/builder $(APP)

build-container-from-dockerfile: hadolint
	docker build -t $(APP) -f Dockerfile .

run-container:
	$(DOCKER) run -d -p $(PORT):$(PORT) --name itotest $(APP)

test:
	go test -v -cover -covermode=atomic ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

stop:
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
