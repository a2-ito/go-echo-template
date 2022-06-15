DOCKER		= docker
DOCKERFILE	= Dockerfile
APP		= itotest
BINARY		= go-echo-template
LINT_IGNORE	= ""

export
PORT	:= 8080

all: format lint test run

lint:
	staticcheck ./...
	golangci-lint run

hadolint:
	$(DOCKER) run --rm -i hadolint/hadolint hadolint - --ignore ${LINT_IGNORE} < $(DOCKERFILE)

format:
	gofmt -l -w .

test:
	go test ./...

run:
	go run main.go

build:
	go build

build-container:
	pack build --builder=gcr.io/buildpacks/builder $(APP)

build-container-from-dockerfile: hadolint
	docker build -t $(APP) -f Dockerfile .

run-container:
	$(DOCKER) run -d -p $(PORT):$(PORT) --name $(APP) $(APP)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: all lint hadolint format test run build build-container build-container-from-dockerfile run-container
